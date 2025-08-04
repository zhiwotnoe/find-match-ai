-- Расширение для векторного поиска
CREATE EXTENSION IF NOT EXISTS vector;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    bio TEXT,
    text_embedding vector(384),  -- Размерность all-MiniLM-L6-v2
    face_embedding vector(512)   -- Размерность FaceNet
);