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
        return response.json()   
    else:
        st.error(f"Erro na requisição: {response.status_code}")
        return None
    

def render_metrics():

    col1, col2 = st.columns([2, 3], gap="large")

    json_data = get_data("http://10.254.19.8:8081/api/v1/location") 


    with col1:
        st.markdown("<h4 style='color: #4CAF50;'>Outputs do Modelo</h4>", unsafe_allow_html=True)

        names = [item['name'] for item in json_data]
        opcao = st.selectbox("Selecione a área", names)
        
        selected_item = next(item for item in json_data if item['name'] == opcao)

        latitude = selected_item['latitude']
        longitude = selected_item['longitude']
        st.write(f"**Latitude:** {latitude}")
        st.write(f"**Longitude:** {longitude}")

        images_bytes = [base64.b64decode(prediction['annotated_image_path']) for prediction in selected_item['predictions']]

        images = [Image.open(BytesIO(img_bytes)) for img_bytes in images_bytes]

        if images:
            st.write("")  

            image_index = st.selectbox("Escolha a imagem", range(len(images)), format_func=lambda x: f"Mês {x + 1}")
            
            detections = selected_item['predictions'][image_index]['detections']
     
            st.write(f"**Quantidade contada:** {detections}")
   
    with col2:
        if images:
            st.image(images[image_index], caption=f"Mês {image_index + 1}", use_column_width=True)
        else:
            st.warning("Não há imagens disponíveis para a área selecionada.")
                

