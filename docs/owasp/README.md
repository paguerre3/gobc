# OWASP Top 10

Here’s a brief summary of the **OWASP Top 10** security risks along with examples, presented in a table for easy reference:

| **Rank** | **Category**                   | **Description**                                                                                         | **Example**                                                                                 |
|----------|--------------------------------|---------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| 1        | **Broken Access Control**      | Weaknesses allowing unauthorized access to restricted resources.                                        | User can modify a URL (`/admin`) to access admin-only data without authorization.           |
| 2        | **Cryptographic Failures**     | Failure to properly encrypt or protect sensitive data.                                                  | Storing **passwords in plaintext** instead of using hashing like bcrypt or Argon2.             |
| 3        | **Injection**                  | **Unsanitized inputs** leading to malicious commands or queries.                                            | SQL Injection: `OR '1'='1'` in a login form bypassing authentication.                      |
| 4        | **Insecure Design**            | Inadequate security controls in the system design phase.                                                | **Lack of rate limiting**, allowing brute-force attacks on login endpoints.                    |
| 5        | **Security Misconfiguration**  | Improper configuration of systems leading to vulnerabilities.                                           | **Running servers with default credentials like `admin:admin`**.                               |
| 6        | **Vulnerable and Outdated Components** | Using components with known vulnerabilities or unsupported versions.                                    | Using an old version of a library like Log4j with known critical exploits.                 |
| 7        | **Identification and Authentication Failures** | Weak or mismanaged identity verification mechanisms.                                                    | Allowing weak passwords or unverified account registration.                                |
| 8        | **Software and Data Integrity Failures** | Failing to validate software updates or data integrity.                                                 | A compromised CI/CD pipeline pushing malicious code to production.                         |
| 9        | **Security Logging and Monitoring Failures** | Insufficient logging or monitoring to detect and respond to breaches.                                   | No alerting mechanism for unusual login attempts or failed access control events.           |
| 10       | **Server-Side Request Forgery (SSRF)** | **Application can be tricked into making malicious requests to unintended locations.**                      | Exploiting an image URL upload feature to access internal servers.                         |

This table provides concise descriptions and real-world examples that can help you articulate each vulnerability.

---
**XSS (Cross-Site Scripting)** is a security vulnerability where attackers inject malicious scripts into web pages viewed by other users. It’s a subset of **Injection** vulnerabilities but focused on client-side scripts (like JavaScript).

Here’s a breakdown for **XSS**:

| **Type**          | **Description**                                                                                      | **Example**                                                                                                           |
|--------------------|------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| **Stored XSS**     | Malicious scripts are permanently stored on the target server (e.g., in a database or message board). | Attacker posts a comment: `<script>alert('Hacked!');</script>` on a forum, executing when viewed by other users.       |
| **Reflected XSS**  | **Malicious scripts** are reflected off a web server in a response, typically **via a URL or input field**.   | A link: `https://example.com/search?q=<script>alert('XSS');</script>` executes the script when clicked by the victim. |
| **DOM-Based XSS**  | The script executes directly in the browser due to insecure handling of DOM elements in JavaScript.   | **A web app uses `document.location` to write input to the DOM without sanitization, enabling script injection.**          |

**Mitigation Techniques:**

1. **Input Validation:** Validate and sanitize all user inputs.
2. **Output Encoding:** **Use libraries like OWASP's *ESAPI* to encode outputs in HTML, JavaScript, or URLs.**
3. **Content Security Policy (CSP):** **Restrict the sources from which scripts can be executed.**
4. **Avoid Inline Scripts:** Use external script files and prevent `eval()` or similar functions.

Would you like additional details or example code for mitigation?

---
### **Man-in-the-Middle (MITM) Attack: Tampering with Data in Transit**

A **Man-in-the-Middle (MITM)** attack occurs when an attacker intercepts and manipulates communications between two parties without their knowledge. This allows the attacker to read, modify, or inject data in transit.

#### **How It Works:**
1. **Interception:** The attacker positions themselves between the victim and the server, intercepting the communication.
2. **Tampering:** The attacker can alter the intercepted data (e.g., change transaction amounts, inject malicious code).
3. **Relay:** The modified data is sent to the intended recipient, making the attack difficult to detect.

#### **Example:**
1. **Scenario:**
   - A user logs into a banking website over an unsecured Wi-Fi network.
   - The attacker intercepts the login request and modifies the HTTP payload to redirect funds to their own account.
2. **Result:**
   - The victim sees a confirmation page, unaware of the tampered data.

#### **Mitigation Techniques:**
1. **Encryption:**
   - Use **HTTPS/TLS** to encrypt communications and ensure integrity.
   - Avoid using insecure protocols like HTTP.
2. **Certificates:**
   - Verify server certificates to prevent spoofed websites or servers.
3. **Secure Networks:**
   - Avoid using public Wi-Fi without a VPN.
   - Use secure, private networks for sensitive transactions.
4. **End-to-End Validation:**
   - Implement cryptographic integrity checks to verify data has not been altered during transit.
