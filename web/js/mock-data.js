// Mock Data for Static Deployment
const MOCK_ITEMS = [
    {
        id: 1,
        name: 'Reindeer Eco',
        type: 'Reindeer',
        category: 'Standard',
        price: 150,
        image_url: 'img/reindeer_eco.png',
        stats: { speed: 10, stamina: 100, capacity: 50, magic: 20 }
    },
    {
        id: 2,
        name: 'Reindeer Performance',
        type: 'Reindeer',
        category: 'Premium',
        price: 300,
        image_url: 'img/reindeer_perf.png',
        stats: { speed: 25, stamina: 90, capacity: 40, magic: 30 }
    },
    {
        id: 3,
        name: 'Reindeer Luxury',
        type: 'Reindeer',
        category: 'Luxury',
        price: 500,
        image_url: 'img/reindeer_lux.png',
        stats: { speed: 15, stamina: 80, capacity: 30, magic: 100 }
    },
    {
        id: 4,
        name: 'Sleigh Eco',
        type: 'Sleigh',
        category: 'Standard',
        price: 200,
        image_url: 'img/sleigh_eco.png',
        stats: { speed: 8, stamina: 0, capacity: 150, magic: 0 }
    },
    {
        id: 5,
        name: 'Sleigh Premium',
        type: 'Sleigh',
        category: 'Premium',
        price: 450,
        image_url: 'img/sleigh_prem.png',
        stats: { speed: 20, stamina: 0, capacity: 120, magic: 50 }
    },
    {
        id: 6,
        name: 'Sleigh Luxury',
        type: 'Sleigh',
        category: 'Luxury',
        price: 800,
        image_url: 'img/sleigh_lux.png',
        stats: { speed: 15, stamina: 0, capacity: 100, magic: 150 }
    },
    {
        id: 7,
        name: 'Rudolph Jr',
        type: 'Reindeer',
        category: 'Special',
        price: 600,
        image_url: 'img/reindeer_rudolph_jr.png',
        stats: { speed: 30, stamina: 110, capacity: 40, magic: 80 }
    },
    {
        id: 8,
        name: 'Blitzen Spark',
        type: 'Reindeer',
        category: 'Special',
        price: 550,
        image_url: 'img/reindeer_blitzen.png',
        stats: { speed: 35, stamina: 85, capacity: 45, magic: 40 }
    },
    {
        id: 9,
        name: 'Cargo Hauler 9000',
        type: 'Sleigh',
        category: 'Heavy',
        price: 700,
        image_url: 'img/sleigh_cargo.png',
        stats: { speed: 10, stamina: 0, capacity: 300, magic: 10 }
    },
    {
        id: 10,
        name: 'Night Fury',
        type: 'Sleigh',
        category: 'Stealth',
        price: 1000,
        image_url: 'img/sleigh_stealth.png',
        stats: { speed: 40, stamina: 0, capacity: 80, magic: 60 }
    }
];

const MOCK_MODS = [
    {
        id: 1,
        name: 'Rocket Booster',
        price: 100,
        stat_boosts: { speed: 20, stamina: -10, capacity: 0, magic: 0 }
    },
    {
        id: 2,
        name: 'Magic Dust sack',
        price: 150,
        stat_boosts: { speed: 5, stamina: 0, capacity: 0, magic: 30 }
    },
    {
        id: 3,
        name: 'Extra Cargo Bay',
        price: 80,
        stat_boosts: { speed: -5, stamina: 0, capacity: 50, magic: 0 }
    },
    {
        id: 4,
        name: 'Titanium Runners',
        price: 120,
        stat_boosts: { speed: 10, stamina: 0, capacity: 20, magic: 0 }
    }
];
