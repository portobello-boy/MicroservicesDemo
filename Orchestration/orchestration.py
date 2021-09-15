from flask import Flask, request, Response
import requests as client
import json

app = Flask(__name__)

@app.route('/enrich/', methods=['POST'])
def enrich():
    resp = Response("")
    resp.headers['content-type'] = 'application/json'

    body = request.json
    if body['type'] == 'getAllEvents':
        resp.data = client.get('http://localhost:3000/events/').text
    elif body['type'] == 'getEvent':
        objectID = body['eventData']['_id']
        resp.data = client.get(f'http://localhost:3000/events/{objectID}').text
    elif body['type'] == 'createEvent':
        eventData = json.dumps(body['eventData'])
        resp.data = client.put(f'http://localhost:3000/events/',data=eventData).text
    elif body['type'] == 'updateEvent':
        eventData = json.dumps(body['eventData'])
        resp.data = client.patch(f'http://localhost:3000/events/',data=eventData).text
    elif body['type'] == 'deleteEvent':
        objectID = body['eventData']['_id']
        resp.data = client.delete(f'http://localhost:3000/events/{objectID}').text
    else:
        resp.status_code = 500
        resp.data = "Invalid request type"

    return resp


if __name__ == '__main__':
    app.run(port=3001)