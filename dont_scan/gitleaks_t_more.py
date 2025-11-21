# gitleaks_test_more.py
# Fresh set of intentionally fake secrets for testing Gitleaks detection.

# 1. New GitHub Machine User Token (ghm_)
GITHUB_MACHINE_TOKEN = "ghm_FAKE1234567890abcdef1234567890abcd"

# 2. New Azure Storage Account Key (common base64 pattern)
AZURE_STORAGE_KEY = "Eby8vdM02xNOcqFlFAKEKEYFAKEKEYFAKEKEYqqDqZx8u2n0="

# 3. New Base64-like API private secret
SERVICE_PRIVATE_SECRET = "MIIEvgIBADANBgkqhkiG9FAKEBASE64SECRET12345678"

def main():
    print("Loaded another fresh set of fake secrets.")

if __name__ == "__main__":
    main()
