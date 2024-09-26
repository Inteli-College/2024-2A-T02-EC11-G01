import streamlit as st
from PIL import Image
import plotly.express as px
import pandas as pd

st.set_page_config(page_title="Visão Geral", layout="wide")

image_path = "../docs/static/img/imagem-modelo-deep-forest.jpeg" 
image = Image.open(image_path)

st.markdown(
    """
    <style>
    @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600&display=swap');
    
    /* Aplica a fonte Poppins globalmente a todos os elementos */
    html, body, h1, h2, h3, h4, h5, h6, p, div, span, li, a, label, button, [class*="css"] {
        font-family: 'Poppins', sans-serif !important;
    }

    /* Estilizando a barra lateral */
    .css-1d391kg {
        background-color: #A6C8B0 !important;
    }

    /* Outros estilos de layout */
    .container {
        background-color: #A6C8B0;
        padding: 20px;
        border-radius: 10px;
    }
    
    .left-menu {
        background-color: #A6C8B0;
        padding: 20px;
        border-radius: 10px;
    }
    
    .card {
        background-color: white;
        border: 1px solid #D3D3D3;
        border-radius: 10px;
        padding: 20px;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
    }
    
    .metric-box {
        background-color: white;
        padding: 20px;
        border-radius: 10px;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
    }
    
    .header {
        text-align: left;
        color: #4CAF50;
    }
    </style>
    """,
    unsafe_allow_html=True
)

with st.sidebar:
    st.image("abundance.png", width=150)
    st.markdown("<h3 style='color: #4CAF50;'>Menu</h3>", unsafe_allow_html=True)
    st.write("Dashboard")
    st.write("Áreas cobertas")
    st.write("Detalhamento")

st.markdown("<h1 class='header'>Dashboard</h1>", unsafe_allow_html=True)
st.markdown("<h6 style='color: #black;'>Visão geral - Métricas</h6>", unsafe_allow_html=True)

col1, col2 = st.columns([3, 2], gap="large")  

with col1:
    st.markdown("<h4 style='color: #4CAF50;'>Mapeamento das áreas</h4>", unsafe_allow_html=True)
    # Métrica 1: Árvores encontradas
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

    # Métrica 2: Tempo de voo
    # st.markdown(
    #     """
    #     <div class='metric-box'>
    #         <h5>Tempo de voo</h5>
    #         <p>De tarde (13h-16h): 1.890 árvores</p>
    #     </div>
    #     """,
    #     unsafe_allow_html=True
    # )

    # # Gráfico de pizza
    # voo_data = pd.DataFrame({
    #     'Período': ['De tarde', 'De noite', 'De manhã'],
    #     'Percentual': [40, 32, 28]
    # })
    # fig_pie = px.pie(voo_data, names='Período', values='Percentual', title="Distribuição do Tempo de Voo",
    #                  color_discrete_sequence=px.colors.sequential.Greens)  
    # fig_pie.update_layout(margin=dict(l=0, r=0, t=100, b=100))  
    # st.plotly_chart(fig_pie, use_container_width=True)

# Coluna 2: Imagem e Detalhes
with col2:
    st.markdown("<h4 style='color: #4CAF50;'>Últimas imagens</h4>", unsafe_allow_html=True)
    
    st.image(image, caption="Imagem retornada pelo modelo de deep forest", use_column_width=True)


st.markdown("<h4 style='color: #4CAF50;'>Quantidade de Carbono Capturado</h4>", unsafe_allow_html=True)
col3, col4 = st.columns([2, 3], gap="large")

with col3: 
    st.metric(
        "Quantidade de Carbono capturado", 
        "2.568 toneladas", 
        "-2.1% vs última semana"
    )

with col4:
    linha_data = pd.DataFrame({
        'Dias': ['1', '2', '3', '4', '5', '6', '7'],
        'Carbono': [1, 2, 3, 4, 5, 6, 7]
    })
    fig_line = px.line(linha_data, x='Dias', y='Carbono', title="Histórico de Carbono Capturado",
                       line_shape='linear', color_discrete_sequence=['#388E3C'])  
    fig_line.update_layout(margin=dict(l=0, r=0, t=100, b=100))  
    st.plotly_chart(fig_line, use_container_width=True)

# Métricas por área monitorada
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
