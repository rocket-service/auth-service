from websocket import create_connection
import json

# Create connection with our web-socket server
ws = create_connection("ws://localhost:3071/handle?m=auth")

# Send JSON styled request to server with user data
ws.send(json.dumps({'k': 'test_key', 'h': 'test_hardware'}))

# Receive result returned from web-socket server
print(f'Received: {ws.recv()}')

# Close our connection
ws.close()