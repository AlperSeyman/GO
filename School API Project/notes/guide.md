1. Marshalling vs. Unmarshalling

    json.Marshal: Go Struct → []byte (JSON).

    json.Unmarshal: []byte (JSON) → Go Struct

2. Encoder vs. Decoder (The Streams)

    json.NewDecoder: Used for reading JSON from a stream (like r.Body in a web request). Create decoder to convert json to struct type

    decoder.Decode: Convert to Go struct type

    json.NewEncoder: Used for writing JSON directly to a stream (like w in an HTTP response). Create encoder to convert struct to json type

    encoder.Encode: Convert to json type

3. Middleware (The Gatekeeper):

    Purpose: To run code before the main logic starts.

    Best Use: Checking if a user is logged in, logging who visited the site, or handling errors.

    Benefit: You don't have to repeat the same "Check" code in every single file.

4. Security Headers

X-Frame-Options: deny

    What it does: It prevents other websites from putting your website inside an <iframe> (a window inside another site).

    Why: This stops "Clickjacking", where a hacker hides your site under a fake button to trick users into clicking something.

X-XSS-Protection: 1;mode=block

    What it does: It tells the browser to look for Cross-Site Scripting (XSS) attacks.

    Why: If the browser sees a suspicious script trying to steal data, it will stop the page from loading entirely.

X-Content-Type-Options: nosniff

    What it does: It tells the browser: "Do not guess the file type. Only use the type I tell you."

    Why: This prevents a hacker from disguising a dangerous script as a harmless image.

Strict-Transport-Security (HSTS)

    What it does: It tells the browser: "Only talk to this server using HTTPS (secure) for the next two years."

    Why: It ensures that even if a user types http://, the browser will automatically switch to the secure https://.

Content-Security-Policy: default-src 'self'

    What it does: This is the most powerful one. It says: "Only trust scripts, images, and data that come from my own server."

    Why: It stops hackers from loading malicious scripts from external, dangerous websites.

Referrer-Policy: no-referrer

    What it does: When a user clicks a link to leave your site, the browser will not tell the new site where the user came from.

    Why: This protects the privacy of your users.

5. CORS Middleware:

    What it does: It gives permission to browsers to let a website talk to your API.

    Where it goes: It wraps around your Router (Mux).

    Key Header: Access-Control-Allow-Origin tells the browser which websites are "friends."