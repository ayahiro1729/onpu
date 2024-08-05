CREATE TABLE musics (
    id SERIAL PRIMARY KEY,
    music_list_id INT NOT NULL REFERENCES music_lists(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    image TEXT,
    artist_name VARCHAR(255),
    spotify_link TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
