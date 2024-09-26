import requests
import streamlit as st

BACKEND_URL = "http://localhost:8080"

def fetch_data(endpoint):
    try:
        response = requests.get(f"{BACKEND_URL}/{endpoint}")
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        st.error(f"Erro ao conectar com o backend: {e}")
        return None
