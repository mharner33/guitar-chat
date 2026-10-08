const RECENT_KEY = "gc_recent";
const CONV_KEY = "gc_conversation";
const RECENT_MAX = 20;

const thread = document.getElementById("thread");
const form = document.getElementById("composer");
const input = document.getElementById("question");
const sendBtn = document.getElementById("send");
const recentList = document.getElementById("recent-list");
const clearBtn = document.getElementById("clear-recent");

function getRecent() {
  try { return JSON.parse(localStorage.getItem(RECENT_KEY)) || []; }
  catch { return []; }
}
function setRecent(list) {
  try { localStorage.setItem(RECENT_KEY, JSON.stringify(list)); }
  catch (err) { console.warn("could not save recent queries:", err); }
}

function getConversationId() {
  try { return localStorage.getItem(CONV_KEY) || null; }
  catch (err) { console.warn("could not read conversation id:", err); return null; }
}
function setConversationId(id) {
  try { localStorage.setItem(CONV_KEY, id); }
  catch (err) { console.warn("could not save conversation id:", err); }
}

function pushRecent(q) {
  let list = getRecent().filter((x) => x !== q);
  list.unshift(q);
  list = list.slice(0, RECENT_MAX);
  setRecent(list);
  renderRecent();
}

function renderRecent() {
  const list = getRecent();
  recentList.innerHTML = "";
  for (const q of list) {
    const li = document.createElement("li");
    const btn = document.createElement("button");
    btn.type = "button";
    btn.textContent = q;
    btn.title = q;
    btn.addEventListener("click", () => { input.value = q; input.focus(); });
    li.appendChild(btn);
    recentList.appendChild(li);
  }
}

function addMessage(role, text, sources) {
  const el = document.createElement("div");
  el.className = "msg " + role;
  el.textContent = text;
  if (sources && sources.length) {
    const wrap = document.createElement("div");
    wrap.className = "sources";
    for (const s of sources) {
      const chip = document.createElement("span");
      chip.className = "chip";
      chip.textContent = `[${s.n}] ${s.title} · ${s.section}`;
      wrap.appendChild(chip);
    }
    el.appendChild(wrap);
  }
  thread.appendChild(el);
  thread.scrollTop = thread.scrollHeight;
}

function setLoading(on) { input.disabled = on; sendBtn.disabled = on; }

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const question = input.value.trim();
  if (!question) return;

  addMessage("user", question);
  pushRecent(question);
  input.value = "";
  setLoading(true);

  try {
    const res = await fetch("/api/chat", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        conversation_id: getConversationId(),
        question,
      }),
    });
    let data = null;
    try { data = await res.json(); }
    catch { data = null; }
    if (!res.ok) {
      addMessage("error", (data && data.error && data.error.message) || ("HTTP " + res.status));
      return;
    }
    if (!data || typeof data !== "object") {
      addMessage("error", "Invalid response from server");
      return;
    }
    if (data.conversation_id) setConversationId(data.conversation_id);
    addMessage("assistant", data.answer, data.sources);
  } catch (err) {
    addMessage("error", "Network error: " + err.message);
  } finally {
    setLoading(false);
    input.focus();
  }
});

clearBtn.addEventListener("click", () => { setRecent([]); renderRecent(); });

renderRecent();
