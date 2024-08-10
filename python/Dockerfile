 FROM python:alpine

 WORKDIR /app

 EXPOSE 5000
 
 COPY . .

 RUN pip install --no-cache-dir -r requirements.txt

 CMD ["python", "app.py"]
