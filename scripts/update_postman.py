import json
from pathlib import Path

path = Path(__file__).resolve().parents[1] / "postman.json"

data = json.loads(path.read_text())

new_payload = {
    "slug": "sample",
    "title": "Sample",
    "shortTitle": "Sample",
    "content": "This is the plain-text content of the blog.",
    "excerpt": "A short excerpt for listing pages.",
    "category": "category2",
    "date": "2026-03-09",
    "author": "Navee",
    "img": "https://example.com/image.png",
    "tags": ["example", "blog"],
    "fullContent": {
        "introduction": "This is an intro section.",
        "sections": [
            {"heading": "Getting started", "content": "Step-by-step details go here."},
            {"heading": "Next steps", "content": "More detailed content in this section."},
        ],
        "quote": "A helpful quote.",
        "conclusion": "This is the conclusion.",
    },
}

for item in data.get("item", []):
    if item.get("name") in ("createblog", "Update"):
        req = item.get("request", {})
        body = req.get("body", {})
        if body.get("mode") == "raw":
            body["raw"] = json.dumps(new_payload, indent=2)

path.write_text(json.dumps(data, indent=2))
print("Updated postman.json payloads for createblog and Update")
