import requests
import json
import argparse
import sys


# Generate self-signed cert for server testing:
# openssl req -x509 -newkey rsa:4096 -keyout secret.key -out secret.crt -sha256 -days 365 -nodes

# Example webhook:
# curl -k -X POST https://<ip_address>:<port>/<route>/<key> -H "Content-Type: application/json" -d '{"data": "test message"}'

def send_hook(url, route, key, data, verify_ssl):
    payload = {"data": data}
    header = {'Content-Type': 'application/json'}
    location = url.rstrip('/') + '/' + route.strip('/') + '/' + key
    if not verify_ssl:
        print("Warning: SSL verification is disabled (verify=False). This is insecure.", file=sys.stderr)
    response = requests.post(
        location,
        data=json.dumps(payload),
        headers=header,
        verify=verify_ssl
    )
    print(response.text)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Send a webhook POST request.")
    parser.add_argument('--url', required=True, help='Base URL of the server')
    parser.add_argument('--route', required=True, help='Route for the action')
    parser.add_argument('--key', required=True, help='Authentication key')
    parser.add_argument('--data', required=True, help='Data to send')
    parser.add_argument('--no-verify', action='store_true', help='Disable SSL verification (insecure)')
    args = parser.parse_args()

    send_hook(args.url, args.route, args.key, args.data, not args.no_verify)