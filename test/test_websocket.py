from websocket import create_connection

# Create connection with our web-socket server
ws = create_connection("ws://localhost:3071/ws/1")

# Send a hello, world string to out web-socket server
ws.send("Hello, World")

# Receive result returned from web-socket server
print(f'Received: {ws.recv()}')

# Close our connection
ws.close()