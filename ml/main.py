from fastapi import FastAPI
from sentence_transformers import SentenceTransformer

app = FastAPI()
model = SentenceTransformer('all-MiniLM-L6-v2')

@app.get("/match")
def match(text1: str, text2: str):
    emb1 = model.encode(text1)
    emb2 = model.encode(text2)
    similarity = (emb1 @ emb2.T) * 100  # Процент совпадения
    return {"match_percent": similarity}