import streamlit as st
import plotly.express as px
import pandas as pd
import base64
from io import BytesIO
from PIL import Image
import requests
import json

def get_data(api_url):
    response = requests.get(api_url)
    if response.status_code == 200:
        return response.json()   # Retorna os dados como um dicionário Python
    else:
        st.error(f"Erro na requisição: {response.status_code}")
        return None
    

def render_metrics():
    col1, col2 = st.columns([3, 2], gap="large")

    json_data = get_data("http://10.128.0.28:8081/api/v1/predictions")  
    primeira_info = json_data[0]
    segunda_info = json_data[1]

    with col1:
        st.markdown("<h4 style='color: #4CAF50;'>Mapeamento das áreas</h4>", unsafe_allow_html=True)
        st.markdown(
            """
            <div class='metric-box' style='margin-bottom: 15px;'>
                <h5>Árvores encontradas</h5>
                <p style='color: green;'>+2.1% vs última semana</p>
            </div>
            """, 
            unsafe_allow_html=True
        )
        
        bar_data = pd.DataFrame({
            'Dias': [primeira_info['created_at'], segunda_info['created_at']],
            'Árvores': [int(primeira_info['detections']), int(segunda_info['detections'])]
        })
        fig_bar = px.bar(bar_data, x='Dias', y='Árvores', title="Árvores encontradas por dia",
                         color_discrete_sequence=['#4CAF50'])
        fig_bar.update_layout(margin=dict(l=0, r=0, t=100, b=100))
        st.plotly_chart(fig_bar, use_container_width=True)

    with col2:

        # json_data = get_data("http://10.128.0.28:8081/api/v1/predictions")  
        # primeira_info = json_data[0]

        st.markdown("<h4 style='color: #4CAF50;'>Outputs do Modelo</h4>", unsafe_allow_html=True)

        img_bytes_annotated = base64.b64decode(primeira_info["annotated_image"])
        img_bytes_raw = base64.b64decode(primeira_info["raw_image"])

        annotated_image = Image.open(BytesIO(img_bytes_annotated))
        raw_image = Image.open(BytesIO(img_bytes_raw))

        imagens = [
            raw_image,
            annotated_image,
        ]

        nomes_imagens = [
            "Original Image",
            "Annotated Image",
        ]


        # Selectbox com as opções de imagem
        opcao = st.selectbox("Contagem realizada", nomes_imagens)

        # Encontrar o índice da imagem correspondente ao nome selecionado
        indice = nomes_imagens.index(opcao)

        if(opcao == "Annotated Image"):
            st.markdown(f"**Quantidade de árvores encontradas:** {primeira_info['detections']}")


        # Exibir a imagem correspondente à opção selecionada
        st.image(imagens[indice], caption=opcao, use_column_width=True)

        # opcao = st.selectbox("Escolha a imagem", [f"Imagem {i+1}" for i in range(len(imagens))])
        # indice = int(opcao.split()[1]) - 1
        # st.image(imagens[indice], caption=opcao, use_column_width=True)

        #sst.image(annotated_image, caption="Imagem em Base64", use_column_width=True)
