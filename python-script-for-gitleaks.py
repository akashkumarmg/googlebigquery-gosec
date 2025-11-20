# gitleaks_test_sample.py
# Contains intentionally fake secrets for testing Gitleaks detection.
# These are NOT real credentials.

import os

# 1. New fake GitHub PAT (different pattern from previous)
GITHUB_TOKEN = "ghu_9xK2lP0Q2FAKEFAKEFAKEFAKE123456"

# 2. New fake GitLab PAT
GITLAB_PAT = "glpat-FAKE1234567890abcdef1234567890"

# 3. New fake AWS keypair (completely new values)
AWS_ACCESS_KEY = "ASIAFAKEKEY999999999"
AWS_SECRET_KEY = "abCDeFGhijklMNOPQRSTUVWXYZ1234567FAKE"

# 4. New fake Stripe Secret Key (Gitleaks detects stripe_)
STRIPE_SECRET_KEY = "sk_live_51FAKEabc123xyz456uvw789"

# 5. New fake Google API Key (common pattern)
GOOGLE_API_KEY = "AIzaSyD-FaKeApIkEy00000000000000000001"

# 6. New fake Twilio API Secret
TWILIO_AUTH_TOKEN = "4f9a8d3e1b2c0fae1234567890abcdef"

# 7. New fake RSA Private Key (shortened but detectable)
RSA_PRIVATE_KEY = """
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCFAKEKEY1234567890ABCDEFGHIJKLMN
OPQRSTUVWXYZ0123456789FAKEKEYabcdefghijklmno
-----END RSA PRIVATE KEY-----
"""

# 8. New fake Slack Bot Token (xoxb- pattern)
SLACK_BOT_TOKEN = "xoxb-9999999999999-8888888888888-FAKEtok3n12345678"

def main():
    print("Loaded fake secrets (new set).")

if __name__ == "__main__":
    main()
