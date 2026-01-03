package handlers

import (
	"fmt"
	"net/http"
	"time"
)

// DocsHandler exibe a documentação interativa
func DocsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	html := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TVMaze API - Documentação Interativa</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 12px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 40px;
            text-align: center;
        }
        .header h1 { font-size: 2.5em; margin-bottom: 10px; }
        .header p { font-size: 1.1em; opacity: 0.9; }
        .content { padding: 40px; }
        .endpoint {
            background: #f8f9fa;
            border-left: 4px solid #667eea;
            border-radius: 8px;
            padding: 20px;
            margin-bottom: 20px;
            transition: transform 0.2s;
        }
        .endpoint:hover { transform: translateX(5px); box-shadow: 0 4px 12px rgba(0,0,0,0.1); }
        .method {
            display: inline-block;
            padding: 5px 15px;
            border-radius: 4px;
            font-weight: bold;
            font-size: 0.9em;
            margin-right: 10px;
        }
        .method.get { background: #28a745; color: white; }
        .endpoint-title { font-size: 1.3em; margin: 10px 0; color: #333; }
        .endpoint-description { color: #666; margin: 10px 0; }
        .endpoint-url {
            background: #2d3748;
            color: #68d391;
            padding: 10px 15px;
            border-radius: 6px;
            font-family: 'Courier New', monospace;
            margin: 10px 0;
            word-break: break-all;
        }
        .try-button {
            background: #667eea;
            color: white;
            border: none;
            padding: 10px 25px;
            border-radius: 6px;
            cursor: pointer;
            font-size: 1em;
            margin-top: 10px;
            transition: background 0.3s;
        }
        .try-button:hover { background: #5568d3; }
        .response {
            background: #1a202c;
            color: #68d391;
            padding: 20px;
            border-radius: 6px;
            margin-top: 15px;
            display: none;
            font-family: 'Courier New', monospace;
            font-size: 0.9em;
            max-height: 400px;
            overflow-y: auto;
            white-space: pre-wrap;
        }
        .response.show { display: block; }
        .loading { color: #fbbf24; }
        .error { color: #ef4444; }
        .params { margin: 15px 0; }
        .param-input {
            padding: 8px 12px;
            border: 2px solid #e2e8f0;
            border-radius: 4px;
            margin: 5px 5px 5px 0;
            font-size: 0.9em;
        }
        .param-label { font-weight: bold; color: #4a5568; margin-right: 10px; }
        .footer {
            background: #f8f9fa;
            padding: 20px;
            text-align: center;
            color: #666;
            border-top: 1px solid #e2e8f0;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>📺 TVMaze API</h1>
            <p>Documentação Interativa - Teste todas as APIs aqui!</p>
            <p style="font-size: 0.9em; margin-top: 10px;">Data: ` + time.Now().Format("02/01/2006") + `</p>
        </div>
        <div class="content">
            <div class="endpoint">
                <span class="method get">GET</span>
                <h3 class="endpoint-title">Informações da API</h3>
                <p class="endpoint-description">Retorna informações gerais sobre a API e endpoints disponíveis</p>
                <div class="endpoint-url">/</div>
                <button class="try-button" onclick="testAPI('/')">🚀 Testar</button>
                <div id="response-home" class="response"></div>
            </div>
            <div class="endpoint">
                <span class="method get">GET</span>
                <h3 class="endpoint-title">Programação de Hoje</h3>
                <p class="endpoint-description">Consulta a programação de TV do dia por país</p>
                <div class="endpoint-url">/schedule?country=<span style="color: #fbbf24;">COUNTRY_CODE</span></div>
                <div class="params">
                    <label class="param-label">País:</label>
                    <input type="text" class="param-input" id="schedule-country" placeholder="US, BR, GB..." value="US">
                </div>
                <button class="try-button" onclick="testSchedule()">🚀 Testar</button>
                <div id="response-schedule" class="response"></div>
            </div>
            <div class="endpoint">
                <span class="method get">GET</span>
                <h3 class="endpoint-title">Buscar Shows</h3>
                <p class="endpoint-description">Busca shows de TV pelo nome</p>
                <div class="endpoint-url">/search?q=<span style="color: #fbbf24;">SHOW_NAME</span></div>
                <div class="params">
                    <label class="param-label">Nome do Show:</label>
                    <input type="text" class="param-input" id="search-query" placeholder="friends, breaking bad..." value="friends">
                </div>
                <button class="try-button" onclick="testSearch()">🚀 Testar</button>
                <div id="response-search" class="response"></div>
            </div>
            <div class="endpoint">
                <span class="method get">GET</span>
                <h3 class="endpoint-title">Detalhes do Show</h3>
                <p class="endpoint-description">Obtém informações detalhadas de um show específico</p>
                <div class="endpoint-url">/show?id=<span style="color: #fbbf24;">SHOW_ID</span></div>
                <div class="params">
                    <label class="param-label">ID do Show:</label>
                    <input type="text" class="param-input" id="show-id" placeholder="431" value="431">
                </div>
                <button class="try-button" onclick="testShow()">🚀 Testar</button>
                <div id="response-show" class="response"></div>
            </div>
        </div>
        <div class="footer">
            <p>💻 Desenvolvido com Go Lang | 📚 Documentação Interativa</p>
            <p style="margin-top: 10px; font-size: 0.9em;">
                <a href="https://github.com/patrickbathu/golang-tvmaze-api" target="_blank" style="color: #667eea; text-decoration: none;">
                    GitHub Repository
                </a>
            </p>
        </div>
    </div>
    <script>
        async function testAPI(endpoint, params = {}) {
            const queryString = new URLSearchParams(params).toString();
            const url = endpoint + (queryString ? '?' + queryString : '');
            const responseId = 'response-' + endpoint.replace('/', 'home');
            const responseEl = document.getElementById(responseId);
            responseEl.className = 'response show loading';
            responseEl.textContent = '⏳ Carregando...';
            try {
                const response = await fetch(url);
                const data = await response.json();
                responseEl.className = 'response show';
                responseEl.textContent = JSON.stringify(data, null, 2);
            } catch (error) {
                responseEl.className = 'response show error';
                responseEl.textContent = '❌ Erro: ' + error.message;
            }
        }
        function testSchedule() {
            const country = document.getElementById('schedule-country').value || 'US';
            testAPICustom('/schedule', { country }, 'response-schedule');
        }
        function testSearch() {
            const query = document.getElementById('search-query').value;
            if (!query) { alert('Por favor, digite o nome de um show'); return; }
            testAPICustom('/search', { q: query }, 'response-search');
        }
        function testShow() {
            const id = document.getElementById('show-id').value;
            if (!id) { alert('Por favor, digite o ID do show'); return; }
            testAPICustom('/show', { id }, 'response-show');
        }
        async function testAPICustom(endpoint, params, responseId) {
            const queryString = new URLSearchParams(params).toString();
            const url = endpoint + '?' + queryString;
            const responseEl = document.getElementById(responseId);
            responseEl.className = 'response show loading';
            responseEl.textContent = '⏳ Carregando...';
            try {
                const response = await fetch(url);
                const data = await response.json();
                responseEl.className = 'response show';
                responseEl.textContent = JSON.stringify(data, null, 2);
            } catch (error) {
                responseEl.className = 'response show error';
                responseEl.textContent = '❌ Erro: ' + error.message;
            }
        }
    </script>
</body>
</html>`
	
	fmt.Fprint(w, html)
}
