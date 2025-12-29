// Shared logic for Santaverse Garage

// Note: 'api' is global from api-client.js

async function fetchItems() {
    return api.getItems();
}

async function fetchGarage() {
    return api.getGarage();
}

async function fetchMods() {
    return api.getMods();
}

async function buyItem(itemId) {
    return api.buyItem(itemId);
}

async function applyMod(vehicleId, modId) {
    return api.applyMod(vehicleId, modId);
}

// Helper to format stats
function formatStats(stats) {
    if (!stats) return '';
    // Handle if stats is JSON string or object
    const s = typeof stats === 'string' ? JSON.parse(stats) : stats;
    let html = '<ul class="stats-list" style="display: grid; grid-template-columns: 1fr 1fr; gap: 0.5rem; list-style: none; padding: 0;">';
    for (const [key, val] of Object.entries(s)) {
        html += `<li style="background: rgba(255,255,255,0.1); padding: 5px 10px; border-radius: 5px; font-size: 0.9rem;">
            <span style="text-transform: capitalize; color: var(--text-dim);">${key}:</span> 
            <strong style="color: var(--accent-light);">${val}</strong>
        </li>`;
    }
    html += '</ul>';
    return html;
}
