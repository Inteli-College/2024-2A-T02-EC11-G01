import streamlit as st
import plotly.express as px
import pandas as pd

def render_dashboard():
    st.markdown("<h4 style='color: #4CAF50;'>Quantidade de Carbono Capturado</h4>", unsafe_allow_html=True)
    
    col3, col4 = st.columns([2, 3], gap="large")
    
    with col3:
        st.metric("Quantidade de Carbono capturado", "2.568 toneladas", "-2.1% vs última semana")
    
    with col4:
        linha_data = pd.DataFrame({
            'Dias': ['1', '2', '3', '4', '5', '6', '7'],
            'Carbono': [1, 2, 3, 4, 5, 6, 7]
        })
        fig_line = px.line(linha_data, x='Dias', y='Carbono', title="Histórico de Carbono Capturado",
                           line_shape='linear', color_discrete_sequence=['#388E3C'])
        fig_line.update_layout(margin=dict(l=0, r=0, t=100, b=100))
        st.plotly_chart(fig_line, use_container_width=True)

    st.markdown("<h4 style='color: #4CAF50;'>Por área monitorada</h4>", unsafe_allow_html=True)
    
    col5, col6 = st.columns(2)
    
    with col5:
        st.markdown(
            """
            <div class='metric-box'>
                <h5>Área 1: Aurora Verde</h5>
                <p>85% saudáveis</p.
            </div>
            """, 
            unsafe_allow_html=True
        )
    with col6:
        st.markdown(
            """
            <div class='metric-box'>
                <h5>Área 2: Mangabeiras</h5>
                <p>92% saudáveis</p.
            </div>
            """, 
            unsafe_allow_html=True
        )
