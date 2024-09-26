import streamlit as st
import plotly.express as px
import pandas as pd

def render_metrics():
    col1, col2 = st.columns([3, 2], gap="large")

    with col1:
        st.markdown("<h4 style='color: #4CAF50;'>Mapeamento das áreas</h4>", unsafe_allow_html=True)
        st.markdown(
            """
            <div class='metric-box'>
                <h5>Árvores encontradas</h5>
                <p style='color: green;'>+2.1% vs última semana</p>
                <p>Árvores de 1-12 Dez, 2023</p>
            </div>
            """, 
            unsafe_allow_html=True
        )
        
        bar_data = pd.DataFrame({
            'Dias': ['1', '2', '3', '4', '5', '6'],
            'Árvores': [10, 15, 20, 18, 12, 14]
        })
        fig_bar = px.bar(bar_data, x='Dias', y='Árvores', title="Árvores encontradas por dia",
                         color_discrete_sequence=['#4CAF50'])
        fig_bar.update_layout(margin=dict(l=0, r=0, t=100, b=100))
        st.plotly_chart(fig_bar, use_container_width=True)

    with col2:
        st.markdown("<h4 style='color: #4CAF50;'>Últimas imagens</h4>", unsafe_allow_html=True)
        st.image("abundance.png", caption="Imagem retornada pelo modelo de deep forest", use_column_width=True)
