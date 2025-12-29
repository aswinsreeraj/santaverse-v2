const API_BASE = '/api';

async function fetchMarketplaceItems() {
    const res = await fetch(`${API_BASE}/marketplace`);
    if (!res.ok) throw new Error('Failed to fetch items');
    return res.json();
}

async function fetchItemDetails(id) {
    const res = await fetch(`${API_BASE}/marketplace/${id}`);
    if (!res.ok) throw new Error('Failed to fetch item');
    return res.json();
}

async function fetchGarage() {
    const res = await fetch(`${API_BASE}/garage`);
    if (!res.ok) throw new Error('Failed to fetch garage');
    return res.json();
}

async function buyItem(itemId) {
    const res = await fetch(`${API_BASE}/buy`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ item_id: parseInt(itemId) })
    });
    if (!res.ok) throw new Error('Failed to buy item');
    return res.json();
}

async function fetchMods() {
    const res = await fetch(`${API_BASE}/mods`);
    if (!res.ok) throw new Error('Failed to fetch mods');
    return res.json();
}

async function applyMod(vehicleId, modId) {
    const res = await fetch(`${API_BASE}/garage/mod`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ vehicle_id: parseInt(vehicleId), mod_id: parseInt(modId) })
    });
    if (!res.ok) throw new Error('Failed to apply mod');
    return res.json();
}

// Helper to format stats
function formatStats(stats) {
    if (!stats) return '';
    // Handle if stats is JSON string or object
    const s = typeof stats === 'string' ? JSON.parse(stats) : stats;
    let html = '<ul class="stats-list">';
    for (const [key, val] of Object.entries(s)) {
        html += `<li><span style="text-transform: capitalize">${key}</span> <strong>${val}</strong></li>`;
    }
    html += '</ul>';
    return html;
}
