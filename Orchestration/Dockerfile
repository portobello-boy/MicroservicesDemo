FROM python

RUN echo "Starting up CRUD server"

COPY . .

RUN pip3 install flask
RUN pip3 install requests

CMD ["orchestration.py"]

ENTRYPOINT ["python3"]

EXPOSE 3001