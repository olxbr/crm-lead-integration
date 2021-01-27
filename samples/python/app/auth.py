from os import getenv
from base64 import b64decode

from fastapi import HTTPException, Header


def verify_token(Authorization: str = Header("")):
    secret_key = f'vivareal:{getenv("SECRET_KEY")}'
    basic_token = Authorization.replace("Basic ", "")
    token = try_to_decode_base64_hash(basic_token)
    if secret_key != token:
        raise HTTPException(
            status_code=401,
            detail="Unauthorized"
        )
    return True


def try_to_decode_base64_hash(str_hash):
    try:
        return b64decode(str_hash).decode('utf-8')
    except Exception:
        return None
