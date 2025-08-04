# NLP + CV-логика
def get_text_match_score(bio1: str, bio2: str) -> float:
    embeddings = model.encode([bio1, bio2])
    return cosine_similarity(embeddings)[0][1]

def get_face_match_score(img1_path: str, img2_path: str) -> float:
    # Используем FaceNet или ArcFace
    return face_model.compare(img1_path, img2_path)