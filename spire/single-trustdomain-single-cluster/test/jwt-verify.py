import os
import json
import subprocess
import jwt
import requests
from jwt import PyJWKClient
from typing import Dict

#token = "<your-token>"
JWT_TOKEN = os.environ['TOKEN']
#issuer = "https://oidc-discovery.apps.tgeer-test.devcluster.openshift.com"

ISSUER = os.environ['ISSUER']
AUDIENCE = os.environ['AUD']
OIDC_JWKS_URI = os.environ['OIDCURL']

"""
# 1. Fetch JWKS URI from discovery document
discovery = requests.get(OIDC_DISCOVERY_URL).json()
jwks_uri = discovery['jwks_uri']

# 2. Get JWKS (public keys)
jwks = requests.get(jwks_uri).json()
keys = {k['kid']: k for k in jwks['keys']}

# 3. Extract KID from JWT header
unverified_header = jwt.get_unverified_header(JWT_TOKEN)
kid = unverified_header['kid']

# 4. Find public key for this KID
if kid not in keys:
    raise Exception(f"KID {kid} not found in JWKS")

public_key = RSAAlgorithm.from_jwk(keys[kid])

# 5. Decode & verify the JWT signature
try:
    decoded = jwt.decode(JWT_TOKEN, public_key, algorithms=['RS256'], audience=AUDIENCE)
    print("✅ JWT Verified. Payload:")
    print(decoded)
except jwt.ExpiredSignatureError:
    print("❌ Token has expired.")
except jwt.InvalidTokenError as e:
    print(f"❌ Invalid token: {e}")

# Monkey-patch requests.get to disable certificate verification
# Backup original requests.get
_original_requests_get = requests.get

# Define insecure version
def insecure_requests_get(*args, **kwargs):
    kwargs['verify'] = False  # Disable TLS cert verification
    return _original_requests_get(*args, **kwargs)

# Monkey patch requests.get globally
requests.get = insecure_requests_get


print (OIDC_JWKS_URI)

# 2. Fetch key using PyJWKClient
jwk_client = PyJWKClient(OIDC_JWKS_URI)
signing_key = jwk_client.get_signing_key_from_jwt(JWT_TOKEN)

# 3. Decode the token
decoded = jwt.decode(
    JWT_TOKEN,
    signing_key.key,
    algorithms=["RS256"],
    audience=AUDIENCE,  # optional
    issuer=ISSUER,
)
print("✅ JWT Verified. Payload:")
print(decoded)
"""
# === Step 1: Fetch JWKS using curl -k (skip TLS verification) ===
try:
    result = subprocess.run(
        ["curl", "-s", "-k", OIDC_JWKS_URI],
        check=True,
        text=True,
        capture_output=True,
    )
    jwks: Dict = json.loads(result.stdout)
except subprocess.CalledProcessError as e:
    raise RuntimeError(f"curl failed: {e.stderr}") from e
except json.JSONDecodeError as e:
    raise ValueError(f"Invalid JSON returned by JWKS endpoint: {e}") from e

# === Step 2: Extract 'kid' from JWT header ===
unverified_header = jwt.get_unverified_header(JWT_TOKEN)
kid = unverified_header.get("kid")
if not kid:
    raise ValueError("No 'kid' found in JWT header")

# === Step 3: Find matching key in JWKS ===
matching_key = next((k for k in jwks["keys"] if k["kid"] == kid), None)
if not matching_key:
    raise ValueError(f"No matching 'kid' found in JWKS for kid={kid}")

# === Step 4: Construct public key from JWKS ===
public_key = jwt.algorithms.RSAAlgorithm.from_jwk(matching_key)

# === Step 5: Verify JWT ===
try:
    payload = jwt.decode(
        JWT_TOKEN,
        key=public_key,
        algorithms=["RS256"],
        audience=AUDIENCE,
    )
    print("✅ JWT is valid. Payload:")
    print(json.dumps(payload, indent=2))
except jwt.ExpiredSignatureError:
    print("❌ Token has expired.")
except jwt.InvalidTokenError as e:
    print(f"❌ Invalid token: {e}")
