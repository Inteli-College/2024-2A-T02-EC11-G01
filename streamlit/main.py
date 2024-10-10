import streamlit as st
from components.sidebar import render_sidebar
from components.dashboard import render_dashboard
from components.metrics import render_metrics

st.set_page_config(page_title="Visão Geral", layout="wide")

with open("style/style.css") as f:
    st.markdown(f"<style>{f.read()}</style>", unsafe_allow_html=True)

render_sidebar()

st.markdown("<h1 class='header'>Dashboard</h1>", unsafe_allow_html=True)
st.markdown("<h6 style='color: #black;'>Visão geral - Métricas</h6>", unsafe_allow_html=True)

render_dashboard()
render_metrics()