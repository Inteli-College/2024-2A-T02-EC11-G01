# Base image
FROM python:3.10.9

# Install system dependencies
RUN apt-get update && apt-get install -y \
  libgl1-mesa-glx \
  libglib2.0-0 \
  libsm6 \
  libxext6 \
  libxrender-dev \
  g++ \
  curl \
  gdal-bin \
  libgdal-dev \
  && apt-get clean

# Set environment variable for GDAL
ENV CPLUS_INCLUDE_PATH=/usr/include/gdal
ENV C_INCLUDE_PATH=/usr/include/gdal
ENV GDAL_VERSION=3.3.0 

# Set the working directory
WORKDIR /app

# Copy requirements file
COPY requirements.txt .

# Install uv
RUN pip install uv

# Install Python dependencies
RUN uv pip install --system --no-cache-dir --no-deps -r requirements.txt -f https://download.pytorch.org/whl/torch_stable.html

# Copy the app files
COPY . .

# Create directories for images
RUN mkdir -p ./images/input ./images/output

# Expose the port FastAPI will run on
EXPOSE 8000

# Start the FastAPI app
CMD ["fastapi", "run", "src/main.py"]
