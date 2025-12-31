import socket
import sys

# Configuration matches your main.arc
HOST = '127.0.0.1'
PORT = 8080

def run_server():
    # Create a TCP/IP socket
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        # Set option to allow the socket to reuse the address immediately after closing
        # This prevents "Address already in use" errors during testing
        s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        
        try:
            s.bind((HOST, PORT))
        except OSError as e:
            print(f"[!] Failed to bind to {HOST}:{PORT}. Is the port already in use? ({e})")
            return

        s.listen()
        print(f"[*] Python TCP Server listening on {HOST}:{PORT}...")
        print("[*] Waiting for Arc client...")

        # Blocking call, waits for a connection
        conn, addr = s.accept()
        
        with conn:
            print(f"[+] Connected by {addr}")
            
            # Receive data (up to 1024 bytes)
            data = conn.recv(1024)
            
            if data:
                # Decode bytes to string
                try:
                    message = data.decode('utf-8').strip()
                    print(f"[<] Received from Arc: \"{message}\"")
                except UnicodeDecodeError:
                    print(f"[<] Received raw bytes: {data}")

                # Send response back to Arc
                response_str = "Greetings from Python!"
                print(f"[>] Sending to Arc:    \"{response_str}\"")
                
                # Encode string to bytes and send
                conn.sendall(response_str.encode('utf-8'))
            else:
                print("[-] Received empty payload.")

        print("[*] Connection closed.")

if __name__ == '__main__':
    try:
        run_server()
    except KeyboardInterrupt:
        print("\n[!] Server stopped.")