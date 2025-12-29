// API Client - Handles Data Access (Mock vs Real)

const IS_MOCK = true; // Set to false to use Real Go Backend

class SantaAPI {
    constructor() {
        this.prefix = '/api';
    }

    async getItems() {
        if (IS_MOCK) return Promise.resolve(MOCK_ITEMS);
        const res = await fetch(`${this.prefix}/items`);
        return res.json();
    }

    async getItem(id) {
        if (IS_MOCK) return Promise.resolve(MOCK_ITEMS.find(i => i.id === id));
        const res = await fetch(`${this.prefix}/items/${id}`);
        return res.json();
    }

    async getGarage() {
        if (IS_MOCK) {
            const owned = JSON.parse(localStorage.getItem('santa_garage')) || [];
            return Promise.resolve(owned);
        }
        const res = await fetch(`${this.prefix}/garage`);
        return res.json();
    }

    async getMods() {
        if (IS_MOCK) return Promise.resolve(MOCK_MODS);
        const res = await fetch(`${this.prefix}/mods`);
        return res.json();
    }

    async buyItem(itemId) {
        if (IS_MOCK) {
            const garage = JSON.parse(localStorage.getItem('santa_garage')) || [];
            const item = MOCK_ITEMS.find(i => i.id === itemId);
            if (!item) throw new Error("Item not found");

            const newVehicle = {
                id: Date.now(), // Generate fake ID
                item_id: item.id,
                item_name: item.name,
                item_type: item.type,
                item_category: item.category,
                // Ensure local paths
                item_image_url: item.image_url.startsWith('/') ? item.image_url.substring(1) : item.image_url,
                current_stats: { ...item.stats },
                has_mods: false
            };
            garage.push(newVehicle);
            localStorage.setItem('santa_garage', JSON.stringify(garage));
            return Promise.resolve({ success: true });
        }

        const res = await fetch(`${this.prefix}/buy`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ item_id: itemId })
        });
        if (!res.ok) throw new Error("Buy failed");
        return res.json();
    }

    async applyMod(vehicleId, modId) {
        if (IS_MOCK) {
            const garage = JSON.parse(localStorage.getItem('santa_garage')) || [];
            const vehicleIdx = garage.findIndex(v => v.id === vehicleId);
            if (vehicleIdx === -1) throw new Error("Vehicle not found");

            const mod = MOCK_MODS.find(m => m.id === modId);
            if (!mod) throw new Error("Mod not found");

            // Apply stats
            const vehicle = garage[vehicleIdx];
            vehicle.current_stats.speed += mod.stat_boosts.speed || 0;
            vehicle.current_stats.stamina += mod.stat_boosts.stamina || 0;
            vehicle.current_stats.capacity += mod.stat_boosts.capacity || 0;
            vehicle.current_stats.magic += mod.stat_boosts.magic || 0;

            vehicle.has_mods = true;
            localStorage.setItem('santa_garage', JSON.stringify(garage));
            return Promise.resolve({ success: true });
        }

        const res = await fetch(`${this.prefix}/garage/mod`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ vehicle_id: vehicleId, mod_id: modId })
        });
        if (!res.ok) throw new Error("Mod failed");
        return res.json();
    }

    async sellVehicle(vehicleId) {
        if (IS_MOCK) {
            let garage = JSON.parse(localStorage.getItem('santa_garage')) || [];
            garage = garage.filter(v => v.id !== vehicleId);
            localStorage.setItem('santa_garage', JSON.stringify(garage));
            return Promise.resolve({ success: true });
        }

        const res = await fetch(`${this.prefix}/garage/sell`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ vehicle_id: vehicleId })
        });
        if (!res.ok) throw new Error("Sell failed");
        return res.json();
    }

    async resetVehicle(vehicleId) {
        if (IS_MOCK) {
            let garage = JSON.parse(localStorage.getItem('santa_garage')) || [];
            const idx = garage.findIndex(v => v.id === vehicleId);
            if (idx === -1) throw new Error("Vehicle not found");

            const item = MOCK_ITEMS.find(i => i.id === garage[idx].item_id);
            if (!item) throw new Error("Original item not found");

            garage[idx].current_stats = { ...item.stats };
            garage[idx].has_mods = false;
            localStorage.setItem('santa_garage', JSON.stringify(garage));
            return Promise.resolve({ success: true });
        }

        const res = await fetch(`${this.prefix}/garage/reset`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ vehicle_id: vehicleId })
        });
        if (!res.ok) throw new Error("Reset failed");
        return res.json();
    }
}

const api = new SantaAPI();
