-- Enable JSONB support (default in modern Postgres, but good to ensure)

DROP TABLE IF EXISTS applied_mods;
DROP TABLE IF EXISTS owned_vehicles;
DROP TABLE IF EXISTS mods;
DROP TABLE IF EXISTS items;

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- 'reindeer' or 'sleigh'
    category VARCHAR(50) NOT NULL, -- 'Economic', 'Premium', 'Luxury', 'Performance'
    price INT NOT NULL,
    image_url TEXT,
    stats JSONB NOT NULL -- { "speed": 10, "stamina": 5, "capacity": 2, "magic": 1 }
);

CREATE TABLE mods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    stat_boosts JSONB NOT NULL -- { "speed": 5, "magic": 2 }
);

CREATE TABLE owned_vehicles (
    id SERIAL PRIMARY KEY,
    item_id INT REFERENCES items(id),
    user_id INT NOT NULL, -- Hardcoded to 1
    current_stats JSONB NOT NULL
);

CREATE TABLE applied_mods (
    id SERIAL PRIMARY KEY,
    owned_vehicle_id INT REFERENCES owned_vehicles(id),
    mod_id INT REFERENCES mods(id)
);

-- Seed Data: Reindeers
INSERT INTO items (name, type, category, price, image_url, stats) VALUES
('Dasher''s Cousin', 'reindeer', 'Economic', 100, '/img/reindeer_eco.png', '{"speed": 5, "stamina": 4, "capacity": 1, "magic": 2}'),
('Comet GT', 'reindeer', 'Performance', 500, '/img/reindeer_perf.png', '{"speed": 9, "stamina": 6, "capacity": 1, "magic": 4}'),
('Vixen Deluxe', 'reindeer', 'Luxury', 800, '/img/reindeer_lux.png', '{"speed": 7, "stamina": 8, "capacity": 2, "magic": 6}'),
('Rudolph Jr', 'reindeer', 'Legendary', 1500, '/img/reindeer_rudolph_jr.png', '{"speed": 6, "stamina": 10, "capacity": 1, "magic": 10}'),
('Blitzen Spark', 'reindeer', 'Performance', 700, '/img/reindeer_blitzen.png', '{"speed": 10, "stamina": 5, "capacity": 1, "magic": 7}');

-- Seed Data: Sleighs
INSERT INTO items (name, type, category, price, image_url, stats) VALUES
('Wooden Classic', 'sleigh', 'Economic', 150, '/img/sleigh_eco.png', '{"speed": 3, "stamina": 10, "capacity": 4, "magic": 0}'),
('Sleigh X-2000', 'sleigh', 'Premium', 600, '/img/sleigh_prem.png', '{"speed": 6, "stamina": 10, "capacity": 6, "magic": 2}'),
('Santa Force One', 'sleigh', 'Luxury', 1200, '/img/sleigh_lux.png', '{"speed": 8, "stamina": 12, "capacity": 10, "magic": 8}'),
('Cargo Hauler 9000', 'sleigh', 'HeavyDuty', 400, '/img/sleigh_cargo.png', '{"speed": 2, "stamina": 15, "capacity": 20, "magic": 1}'),
('Night Fury', 'sleigh', 'Stealth', 2000, '/img/sleigh_stealth.png', '{"speed": 10, "stamina": 8, "capacity": 2, "magic": 5}');

-- Seed Data: Mods
INSERT INTO mods (name, price, stat_boosts) VALUES
('Turbo Magic Booster', 200, '{"speed": 5, "magic": -1}'),
('Glow-Nose Upgrade', 150, '{"magic": 5, "stamina": 1}'),
('Weather Resistance Kit', 100, '{"stamina": 5}'),
('Gift Expander 3000', 300, '{"capacity": 5, "speed": -2}');
