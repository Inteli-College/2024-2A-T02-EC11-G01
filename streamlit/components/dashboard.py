import streamlit as st
import plotly.express as px
import pandas as pd
import requests

def render_dashboard():
    st.markdown("<h4 style='color: #4CAF50; margin-top: 35px;' >Quantidade de Carbono Capturado</h4>", unsafe_allow_html=True)
    
    col3, col4 = st.columns([2, 3], gap="large")
    
    with col3:
        st.metric("Quantidade de Carbono capturado", "568 toneladas")
    
    with col4:
        linha_data = pd.DataFrame({
            'Dias': ['1', '2', '3', '4', '5'],
            'Carbono': [128, 245 , 230, 323, 568]
        })
        fig_line = px.line(linha_data, x='Dias', y='Carbono',
                           title="Histórico de Carbono Capturado",
                           line_shape='linear', color_discrete_sequence=['#388E3C'])
        fig_line.update_layout(
            title=dict(text="Histórico de Carbono Capturado", font=dict(family="Poppins, sans-serif", size=20)),
            margin=dict(l=0, r=0, t=100, b=100),
            font=dict(family="Poppins, sans-serif") 
        )
        st.plotly_chart(fig_line, use_container_width=True)
    


if __name__ == "__main__":
    render_dashboard()
