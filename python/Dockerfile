 FROM python:alpine

 WORKDIR /app
 
 COPY . .

 RUN pip install --no-cache-dir -r requirements.txt

 EXPOSE 5000

 ENTRYPOINT [ "python", "/app/app.py" ]
