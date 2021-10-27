from flask import Flask, request, Response
import requests as client
import json

app = Flask(__name__)

@app.get('/health')
def health():
    resp = Response('Healthy',status=200)
    resp.headers['Access-Control-Allow-Origin'] = '*'
    print('Healthy')
    return resp

def enrich_options(resp):
    resp.headers['Access-Control-Allow-Headers'] = 'Content-Type'
    resp.headers['Access-Control-Allow-Methods'] = 'POST, OPTIONS'
    resp.headers['Access-Control-Allow-Origin'] = '*'
    resp.data = 'OPTIONS, POST'
    return resp

@app.route('/enrich/', methods=['POST','OPTIONS'])
def enrich():
    resp = Response("")
    resp.headers['found'] = 'true'
    print(request.method,request.data,request.headers)
    if request.method == 'OPTIONS':
        return enrich_options(resp)
    resp.headers['content-type'] = 'application/json'
    # resp.headers['Access-Control-Allow-Headers'] = 'Content-Type'
    # resp.headers['Access-Control-Allow-Methods'] = 'GET, POST, OPTIONS'
    resp.headers['Access-Control-Allow-Origin'] = '*'

    body = request.json
    if body['type'] == 'getAllEvents':
        events = json.loads(client.get('http://localhost:3000/events/').text)
        for e in events:
            e['enriched'] = True
        resp.data = json.dumps(events)
        
    elif body['type'] == 'getEvent':
        objectID = body['eventData']['_id']
        event = json.loads(client.get(f'http://localhost:3000/events/{objectID}').text)
        event['enriched'] = True
        resp.data = json.dumps(event)

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
    app.run(host='0.0.0.0', port=3001)