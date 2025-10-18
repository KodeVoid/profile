package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kendrick | Backend Developer</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
            background: #0d1117;
            color: #c9d1d9;
            line-height: 1.6;
            padding: 2rem;
        }
        
        .container {
            max-width: 900px;
            margin: 0 auto;
            background: #161b22;
            border: 1px solid #30363d;
            border-radius: 6px;
            padding: 3rem;
        }
        
        .header {
            border-bottom: 1px solid #30363d;
            padding-bottom: 2rem;
            margin-bottom: 2rem;
        }
        
        .header h1 {
            font-size: 2.5rem;
            color: #58a6ff;
            margin-bottom: 0.5rem;
        }
        
        .tagline {
            font-size: 1.1rem;
            color: #8b949e;
            margin-bottom: 1.5rem;
        }
        
        .intro {
            color: #c9d1d9;
            line-height: 1.8;
        }
        
        h2 {
            font-size: 1.8rem;
            color: #58a6ff;
            margin: 2rem 0 1rem 0;
            padding-bottom: 0.3rem;
            border-bottom: 1px solid #30363d;
        }
        
        .emoji {
            margin-right: 0.5rem;
        }
        
        .section {
            margin-bottom: 2rem;
        }
        
        ul {
            list-style: none;
            padding-left: 0;
        }
        
        ul li {
            padding: 0.5rem 0;
            color: #c9d1d9;
        }
        
        ul li strong {
            color: #58a6ff;
            display: inline-block;
            min-width: 150px;
        }
        
        .tech-list {
            color: #8b949e;
        }
        
        .stats {
            display: flex;
            gap: 1rem;
            margin: 1.5rem 0;
            flex-wrap: wrap;
        }
        
        .stat-img {
            max-width: 100%;
            border-radius: 6px;
            border: 1px solid #30363d;
        }
        
        .contact-links {
            display: flex;
            gap: 1.5rem;
            margin-top: 1.5rem;
            flex-wrap: wrap;
        }
        
        .contact-links a {
            color: #58a6ff;
            text-decoration: none;
            padding: 0.5rem 1rem;
            border: 1px solid #30363d;
            border-radius: 6px;
            transition: all 0.3s ease;
            display: inline-block;
        }
        
        .contact-links a:hover {
            background: #30363d;
            border-color: #58a6ff;
        }
        
        .api-section {
            background: #0d1117;
            border: 1px solid #30363d;
            border-radius: 6px;
            padding: 1.5rem;
            margin-top: 2rem;
        }
        
        .api-section h3 {
            color: #58a6ff;
            margin-bottom: 1rem;
        }
        
        .endpoint {
            display: flex;
            align-items: center;
            margin: 0.8rem 0;
            font-family: 'Courier New', monospace;
            font-size: 0.95rem;
        }
        
        .method {
            background: #238636;
            color: #fff;
            padding: 0.3rem 0.8rem;
            border-radius: 4px;
            font-weight: bold;
            margin-right: 1rem;
            font-size: 0.85rem;
        }
        
        .path {
            color: #c9d1d9;
        }
        
        .path a {
            color: #58a6ff;
            text-decoration: none;
        }
        
        .path a:hover {
            text-decoration: underline;
        }
        
        .quote {
            text-align: center;
            margin-top: 3rem;
            padding-top: 2rem;
            border-top: 1px solid #30363d;
            color: #8b949e;
            font-style: italic;
        }
        
        @media (max-width: 768px) {
            body {
                padding: 1rem;
            }
            
            .container {
                padding: 1.5rem;
            }
            
            .header h1 {
                font-size: 2rem;
            }
            
            .stats {
                flex-direction: column;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Hey there, I'm Kendrick 👋</h1>
            <p class="tagline">Backend Developer | Automation Architect | Performance Optimizer</p>
            <p class="intro">I craft <strong>high-performance backend systems</strong> and eliminate repetitive work through <strong>intelligent automation</strong>. My passion lies in building software that scales gracefully and workflows that run themselves.</p>
        </div>
        
        <div class="section">
            <h2><span class="emoji">🚀</span>What I Do</h2>
            <ul>
                <li><strong>Backend Engineering:</strong> Design and build robust REST APIs, microservices, and distributed systems that handle real-world scale</li>
                <li><strong>Automation & Integration:</strong> Transform manual processes into automated workflows using n8n, custom scripts, and smart integrations</li>
                <li><strong>Performance Optimization:</strong> Leverage resource-efficient languages like Go and Rust to build systems that do more with less</li>
                <li><strong>Full-Stack When Needed:</strong> Integrate React frontends to complete the user experience</li>
            </ul>
        </div>
        
        <div class="section">
            <h2><span class="emoji">🛠️</span>Tech Arsenal</h2>
            <ul>
                <li><strong>Languages:</strong> <span class="tech-list">Go • Rust • Python • JavaScript</span></li>
                <li><strong>Backend:</strong> <span class="tech-list">REST APIs • Microservices • Distributed Systems</span></li>
                <li><strong>Automation:</strong> <span class="tech-list">n8n • Shell Scripting • Workflow Orchestration</span></li>
                <li><strong>Infrastructure:</strong> <span class="tech-list">Docker • Git • GitHub Actions • CI/CD</span></li>
                <li><strong>Frontend:</strong> <span class="tech-list">React • Modern Web Standards</span></li>
            </ul>
        </div>
        
        <div class="section">
            <h2><span class="emoji">💡</span>Featured Work</h2>
            <p>🚧 <strong>Currently building something exciting...</strong> Check back soon for updates!</p>
            <p style="margin-top: 1rem; color: #8b949e;">In the meantime, explore my repositories to see what I'm working on.</p>
        </div>
        
        <div class="section">
            <h2><span class="emoji">📊</span>GitHub Activity</h2>
            <div class="stats">
                <img class="stat-img" src="https://github-readme-stats.vercel.app/api?username=KodeVoid&show_icons=true&theme=tokyonight&hide_border=true&bg_color=0D1117" alt="Kendrick's GitHub Stats" height="165"/>
                <img class="stat-img" src="https://github-readme-stats.vercel.app/api/top-langs/?username=KodeVoid&layout=compact&theme=tokyonight&hide_border=true&bg_color=0D1117" alt="Top Languages" height="165"/>
            </div>
            <div style="margin-top: 1rem;">
                <img class="stat-img" src="https://github-readme-streak-stats.herokuapp.com/?user=KodeVoid&theme=tokyonight&hide_border=true&background=0D1117" alt="GitHub Streak" height="165"/>
            </div>
        </div>
        
        <div class="api-section">
            <h3>📡 API Endpoints</h3>
            <div class="endpoint">
                <span class="method">GET</span>
                <span class="path"><a href="/me">/me</a> - Get profile information with random cat fact</span>
            </div>
            <div class="endpoint">
                <span class="method">GET</span>
                <span class="path">/health - Service health check</span>
            </div>
        </div>
        
        <div class="section">
            <h2><span class="emoji">🤝</span>Let's Connect</h2>
            <p>I'm always interested in discussing <strong>distributed systems</strong>, <strong>automation strategies</strong>, and <strong>challenging backend problems</strong>. Open to collaboration on impactful projects.</p>
            <div class="contact-links">
                <a href="https://linkedin.com/in/kodevoid" target="_blank">LinkedIn</a>
                <a href="https://x.com/olorikendrick" target="_blank">X / Twitter</a>
                <a href="https://void-iota-ashen.vercel.app" target="_blank">Portfolio</a>
                <a href="https://github.com/KodeVoid" target="_blank">GitHub</a>
            </div>
        </div>
        
        <div class="quote">
            💻 "Automate the boring stuff, architect the interesting stuff."
        </div>
    </div>
</body>
</html>`

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(html))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}

}
