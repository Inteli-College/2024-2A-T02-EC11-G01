FROM python:3.10.13

RUN apt-get update && apt-get install -y \
    gdal-bin \
    libgdal-dev \
    libgl1-mesa-glx \
    python3-dev \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

ENV CPLUS_INCLUDE_PATH=/usr/include/gdal
ENV C_INCLUDE_PATH=/usr/include/gdal

WORKDIR /serverless

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["fastapi", "run", "main.py"]
