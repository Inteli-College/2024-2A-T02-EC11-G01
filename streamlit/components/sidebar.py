import streamlit as st

def render_sidebar():
    with st.sidebar:
        st.image("abundance.png", width=150)
        st.markdown("<h3 style='color: #4CAF50;'>Menu</h3>", unsafe_allow_html=True)
        st.write("Dashboard")
        st.write("Áreas cobertas")
        st.write("Detalhamento")
