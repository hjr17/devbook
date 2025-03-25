package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario renderiza a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "register.html", nil)
}

// CarregarPaginaPrincipal carrega a página principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
	}

	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []models.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// CarregarPaginaDeAtualizacaoDePublicacao carrega a página de edição de publicação
func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)
	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacao models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

// CarregarPaginaDeUsuarios carrega a página com os usuários que atendem o filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick)

	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var usuarios []models.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario carrega a página do perfil do usuário
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", 302)
		return
	}

	usuario, erro := models.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         models.Usuario
		usuarioLogadoID uint64
	}{
		Usuario:         usuario,
		usuarioLogadoID: usuarioLogadoID,
	})
}

// CarregarPerfilDoUsuarioLogado carrega a página do perfil do usuário logado
func CarregarPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := models.BuscarUsuarioCompleto(usuarioID, r)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)
}
