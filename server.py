import http.server
import socketserver

# Set the port you want to listen on
PORT = 8000

# Create a request handler that serves files from the current directory
Handler = http.server.SimpleHTTPRequestHandler

# Set up the server
with socketserver.TCPServer(("", PORT), Handler) as httpd:
    print(f"Serving at port {PORT}")
    # Start the server
    httpd.serve_forever()
