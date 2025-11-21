# gitleaks_test_more.py
# Fresh set of intentionally fake secrets for testing Gitleaks detection.

# 1. New GitHub Machine User Token (ghm_)
GITHUB_MACHINE_TOKEN = "ghm_FAKE1234567890abcdef1234567890abcd"

# 2. New Azure Storage Account Key (common base64 pattern)
AZURE_STORAGE_KEY = "Eby8vdM02xNOcqFlFAKEKEYFAKEKEYFAKEKEYqqDqZx8u2n0="

# 3. New Heroku API Key (heroku_ pattern)
HEROKU_API_KEY = "heroku_api_key_FAKE1234567890abcdef123456"

# 4. New Mailchimp API Key (ends with region)
MAILCHIMP_KEY = "1234567890abcdef1234567890abcdef-us99"

# 5. New JWT Signing Secret (high-entropy base64)
JWT_SIGNING_SECRET = "ZXhhbXBsZS1qd3Qtc2lnbmluZy1zZWNyZXQtRkFLRS1LRVk="

# 6. New PostgreSQL connection string with password
POSTGRES_URL = "postgres://admin:SuperSecretPass123!@localhost:5432/mydb"

# 7. New Docker Hub Access Token (dhp_ fake)
DOCKER_HUB_TOKEN = "dhp_FAKE1234567890abcdef1234567890abcd"

# 8. New Base64-like API private secret
SERVICE_PRIVATE_SECRET = "MIIEvgIBADANBgkqhkiG9FAKEBASE64SECRET12345678"

def main():
    print("Loaded another fresh set of fake secrets.")

if __name__ == "__main__":
    main()
