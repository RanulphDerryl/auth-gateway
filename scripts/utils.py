# utils.py
from typing import List, Dict
import logging
from fastapi import HTTPException

logger = logging.getLogger("auth_gateway")

def get_user_id_from_token(token: str) -> str:
    """
    Extracts and returns the user ID from a given token.
    
    Args:
    - token (str): The token to extract the user ID from.
    
    Returns:
    - str: The extracted user ID.
    """
    # For demonstration purposes, the token is assumed to be a JSON Web Token (JWT) with a user ID in the payload.
    # In a real scenario, you would use a library like PyJWT to handle JWTs.
    import json
    token_data = json.loads(token)
    return token_data.get("user_id")

def validate_user_data(user_data: Dict[str, str]) -> bool:
    """
    Validates the provided user data.
    
    Args:
    - user_data (Dict[str, str]): The user data to validate.
    
    Returns:
    - bool: True if the user data is valid, False otherwise.
    """
    required_fields = ["username", "email", "password"]
    for field in required_fields:
        if field not in user_data:
            logger.error(f"Missing required field: {field}")
            return False
    return True

def handle_http_exception(exception: HTTPException) -> Dict[str, str]:
    """
    Handles HTTP exceptions and returns a standardized error response.
    
    Args:
    - exception (HTTPException): The HTTP exception to handle.
    
    Returns:
    - Dict[str, str]: A standardized error response containing the HTTP status code and error message.
    """
    return {
        "error": str(exception.detail),
        "status_code": exception.status_code,
    }