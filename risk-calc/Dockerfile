FROM python:3.13-alpine
LABEL maintainer="mukashevasset" description="CLI-калькулятор кибер-рисков"
WORKDIR /app
COPY risk_calc.py .
COPY data/ data/
RUN adduser -D app
USER app
ENTRYPOINT ["python", "/app/risk_calc.py"]
