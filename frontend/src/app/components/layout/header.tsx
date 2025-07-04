import React from 'react';

const Header: React.FC = () => {
  return (
    <div id="nav-belt" className="w-full relative z-30" style={{ top: '0px' }}>
      <div className="flex items-center justify-between h-16 px-4 bg-[#131921] text-white">
        {/* Left Section */}
        <div className="flex items-center space-x-4">
          {/* Logo */}
          <div
            id="nav-logo"
            className="flex-shrink-0 flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-14 px-2"
          >
            <a
              href="/ref=nav_logo"
              id="nav-logo-sprites"
              className="flex items-center h-8 relative"
              aria-label="Amazon.com.br"
              lang="en"
            >
              <span className="bg-white w-20 h-8 block rounded-sm mr-1"></span>
              <span id="logo-ext" className="text-sm font-semibold text-white">.com.br</span>
            </a>
          </div>


          {/* Global Location */}
          <div id="nav-global-location-slot" className="ml-4 flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-14 px-2">
            <span
              id="nav-global-location-data-modal-action"
              className="cursor-pointer text-gray-300 hover:text-white"
              data-a-modal="{&quot;width&quot;:375, &quot;closeButton&quot;:&quot;true&quot;,&quot;popoverLabel&quot;:&quot;Escolha sua localização&quot;, &quot;ajaxHeaders&quot;:{&quot;anti-csrftoken-a2z&quot;:&quot;hHnytIHcn69D2ZKOJn4nJ8E/6qFPfvx4S+3nYlrs08TAAAAAAGhlhX8AAAAB&quot;}, &quot;name&quot;:&quot;glow-modal&quot;, &quot;url&quot;:&quot;/portal-migration/hz/glow/get-rendered-address-selections?deviceType=desktop&amp;pageType=Gateway&amp;storeContext=NoStoreName&amp;actionSource=desktop-modal&quot;, &quot;footer&quot;:null,&quot;header&quot;:&quot;Escolha sua localização&quot;}"
              data-action="a-modal"
            >
              <a id="nav-global-location-popover-link" role="button" tabIndex={0} className="flex items-center text-sm" href="">
                {/* Placeholder for nav-packard-glow-loc-icon */}
                <div id="nav-packard-glow-loc-icon" className="w-4 h-4 bg-gray-400 mr-1 rounded-full"></div>
                <div id="glow-ingress-block">
                  <span className="block text-gray-400 text-xs" id="glow-ingress-line1">
                    A entrega será feita em Frederico... 98400000
                  </span>
                  <span className="block text-white font-bold text-sm" id="glow-ingress-line2">
                    Atualizar CEP
                  </span>
                </div>
              </a>
            </span>
            <input id="unifiedLocation1ClickAddress" name="dropdown-selection" type="hidden" value="add-new" />
            <input id="ubbShipTo" name="dropdown-selection-ubb" type="hidden" value="add-new" />
            <input id="glowValidationToken" name="glow-validation-token" type="hidden" value="hHnytIHcn69D2ZKOJn4nJ8E/6qFPfvx4S+3nYlrs08TAAAAAAGhlhX8AAAAB" />
            <input id="glowDestinationType" name="glow-destination-type" type="hidden" value="IP2LOCATION" />
          </div>
        </div>

        {/* Search Bar */}
        <div id="nav-fill-search" className="flex-grow flex justify-center mx-4">
          <div id="nav-search" className="relative w-full max-w-2xl">
            <form id="nav-search-bar-form" acceptCharset="utf-8" action="/s/ref=nb_sb_noss" className="flex items-center justify-between rounded-md overflow-hidden group" method="GET" name="site-search" role="search">
              <input type="hidden" name="__mk_pt_BR" value="ÅMÅŽÕÑ" />

              {/* Wrapper para alinhar altura dos elementos */}
              <div className="flex items-center h-11 w-full bg-white group-focus-within:border-2 group-focus-within:border-[#FF9900]">
                {/* Dropdown */}
                <div className="flex-shrink-0 h-10 flex items-center">
                  <div id="nav-search-dropdown-card" className="relative h-10 flex items-center">
                    <div className="bg-gray-200 text-gray-800 px-3 py-2 text-sm cursor-pointer rounded-l-md flex items-center h-10">
                      <span id="nav-search-label-id" className="whitespace-nowrap">Todos</span>
                      <span>
                      <svg
                        className="text-[#666666] w-4 h-4"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        data-t="dropdown-svg"
                        aria-hidden="true"
                        role="img"
                        fill='currentColor'
                        ><path d="M7 10h10l-5 5z"></path>
                        </svg>
                      </span>
                    </div>
                    <label id="searchDropdownDescription" htmlFor="searchDropdownBox" className="sr-only">Selecione o departamento que deseja pesquisar no</label>
                    <select
                      aria-describedby="searchDropdownDescription"
                      className="absolute top-0 left-0 w-full h-full opacity-0 cursor-pointer text-black"
                      id="searchDropdownBox"
                      name="url"
                      tabIndex={0}
                      title="Pesquisar em"
                    >
                      <option selected value="search-alias=aps">Todos os departamentos</option>
                      <option value="search-alias=alexa-skills">Alexa Skills</option>
                      <option value="search-alias=grocery">Alimentos e Bebidas</option>
                      <option value="search-alias=warehouse-deals">Amazon Quase Novo</option>
                      <option value="search-alias=mobile-apps">Apps e Jogos</option>
                      <option value="search-alias=audible">Audiolivros Audible</option>
                      <option value="search-alias=automotive">Automotivo</option>
                      <option value="search-alias=baby">Bebês</option>
                      <option value="search-alias=beauty">Beleza</option>
                      <option value="search-alias=luxury-beauty">Beleza de Luxo</option>
                      <option value="search-alias=fashion-luggage">Bolsas, Malas e Mochilas</option>
                      <option value="search-alias=toys">Brinquedos e Jogos</option>
                      <option value="search-alias=home">Casa</option>
                      <option value="search-alias=popular">CD e Vinil</option>
                      <option value="search-alias=computers">Computadores e Informática</option>
                      <option value="search-alias=kitchen">Cozinha</option>
                      <option value="search-alias=amazon-devices">Dispositivos Amazon</option>
                      <option value="search-alias=dvd">DVD e Blu-Ray</option>
                      <option value="search-alias=appliances">Eletrodomésticos</option>
                      <option value="search-alias=electronics">Eletrônicos</option>
                      <option value="search-alias=sporting">Esportes e Aventura</option>
                      <option value="search-alias=hi">Ferramentas e Materiais de Construção</option>
                      <option value="search-alias=videogames">Games</option>
                      <option value="search-alias=mi">Instrumentos Musicais</option>
                      <option value="search-alias=garden">Jardim e Piscina</option>
                      <option value="search-alias=stripbooks">Livros</option>
                      <option value="search-alias=digital-text">Loja Kindle</option>
                      <option value="search-alias=office-products">Material para Escritório e Papelaria</option>
                      <option value="search-alias=furniture">Móveis e Decoração</option>
                      <option value="search-alias=pets">Pet Shop</option>
                      <option value="search-alias=instant-video">Prime Video</option>
                      <option value="search-alias=industrial">Produtos Industriais e Científicos</option>
                      <option value="search-alias=specialty-aps-sns">Programe e Poupe</option>
                      <option value="search-alias=fashion">Roupas, Calçados e Joias</option>
                      <option value="search-alias=fashion-womens">Roupas, sapatos e joias femininas</option>
                      <option value="search-alias=fashion-mens">Roupas, sapatos e joias masculinas</option>
                      <option value="search-alias=fashion-girls">Roupas, sapatos e joias para meninas</option>
                      <option value="search-alias=fashion-boys">Roupas, sapatos e joias para meninos</option>
                      <option value="search-alias=fashion-baby">Roupas, sapatos e joias para bebês</option>
                      <option value="search-alias=hpc">Saúde e Cuidados Pessoais</option>
                    </select>
                  </div>
                </div>
                {/* Input */}
                <div className="flex-grow h-10 flex items-center">
                  <div className="relative w-full h-10 flex items-center">
                    <label htmlFor="twotabsearchtextbox" className="sr-only">Pesquisar Amazon.com.br</label>
                    <input
                      type="text"
                      id="twotabsearchtextbox"
                      value=""
                      name="field-keywords"
                      autoComplete="off"
                      placeholder="Pesquisar Amazon.com.br"
                      className="w-full h-10 px-3 py-2 text-gray-900 focus:outline-none focus-none"
                      dir="auto"
                      tabIndex={0}
                      aria-label="Pesquisar Amazon.com.br"
                      role="searchbox"
                      aria-autocomplete="list"
                      aria-controls="sac-autocomplete-results-container"
                      aria-expanded="false"
                      aria-haspopup="grid"
                      spellCheck="false"
                    />
                  </div>
                  <div id="nav-iss-attach"></div>
                </div>
                {/* Botão de busca */}
                <div className="flex-shrink-0 flex items-center h-10">
                  <div className="bg-[#FEBD69] hover:bg-orange-500 h-12 w-10 flex items-center justify-center rounded-r-md cursor-pointer">
                    <span id="nav-search-submit-text" aria-label="Ir">
                      <input id="nav-search-submit-button" type="submit" className="sr-only" value="Ir" tabIndex={0} />
                      <svg 
                          className="text-[#333333] w-6 h-6" 
                          xmlns="http://www.w3.org/2000/svg" 
                          viewBox="0 0 24 24" 
                          data-t="search-svg" 
                          aria-hidden="true" 
                          role="img"
                          fill='currentColor'
                          >
                          <path fillRule="evenodd" clipRule="evenodd" d="M10.5 19C12.4879 19 14.3164 18.3176 15.7641 17.1742L21.2927 22.7069L22.7074 21.2931L17.1778 15.7595C18.319 14.3126 19 12.4858 19 10.5C19 5.80558 15.1944 2 10.5 2C5.80558 2 2 5.80558 2 10.5C2 15.1944 5.80558 19 10.5 19ZM10.5 17C14.0899 17 17 14.0899 17 10.5C17 6.91015 14.0899 4 10.5 4C6.91015 4 4 6.91015 4 10.5C4 14.0899 6.91015 17 10.5 17Z">
                          </path></svg>
                    </span>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>

        {/* Right Section */}
        <div className="flex items-center space-x-6">
          {/* Account & Lists */}
          <div id="nav-link-accountList" className="group relative cursor-pointer flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-14 px-2">
            <a
              href="https://www.amazon.com.br/ap/signin?openid.pape.max_auth_age=0&amp;openid.return_to=https%3A%2F%2Fwww.amazon.com.br%2F%3F%26tag%3Dhydrbrabk-20%26ref%3Dnav_ya_signin%26adgrpid%3D155790195778%26hvpone%3D%26hvptwo%3D%26hvadid%3D677606588104%26hvpos%3D%26hvnetw%3Dg%26hvrand%3D10665990402867838904%26hvqmt%3De%26hvdev%3Dc%26hvdvcmdl%3D%26hvlocint%3D%26hvlocphy%3D9221782%26hvtargid%3Dkwd-10573980%26hydadcr%3D26346_11691057%26gad_source%3D1&amp;openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&amp;openid.assoc_handle=brflex&amp;openid.mode=checkid_setup&amp;openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&amp;openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0"
              className="text-white text-sm"
              data-nav-ref="nav_ya_signin"
              data-nav-role="signin"
              tabIndex={0}
              aria-controls="nav-flyout-accountList"
              data-csa-c-id="h7tsw5-5l7ntd-6ct2tk-9b7gi0"
            >
              <div className="leading-tight">
                <span id="nav-link-accountList-nav-line-1" className="block text-xs text-gray-300 group-hover:text-white">Olá, faça seu login</span>
                <span className="block font-bold">Contas e Listas <i className="ml-1 fas fa-caret-down text-xs"></i></span>
              </div>
            </a>
            <button className="sr-only" aria-label="Expandir conta e listas" tabIndex={0}></button>
          </div>

          {/* Returns & Orders */}
          <a href="/gp/css/order-history?ref_=nav_orders_first" className="text-white text-sm leading-tight flex flex-col items-center justify-center hover:outline-2 hover:outline-white transition-all h-14 px-2" id="nav-orders" tabIndex={0}>
            <span className="block text-xs text-gray-300">Devoluções</span>
            <span className="block font-bold">e Pedidos</span>
          </a>

          {/* Cart */}
          <a href="/gp/cart/view.html?ref_=nav_cart" aria-label="0 itens no carrinho" className="relative flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-14 px-2" id="nav-cart">
            <div id="nav-cart-count-container" className="relative">
              <span id="nav-cart-count" aria-hidden="true" className="absolute -top-2 left-3 text-orange-400 font-bold text-sm">0</span>
              {/* Placeholder for nav-cart-icon */}
              <div className="w-8 h-8 bg-gray-700 rounded-full flex items-center justify-center">
                <i className="fas fa-shopping-cart text-white text-lg"></i> {/* Using Font Awesome for cart icon */}
              </div>
            </div>
            <div id="nav-cart-text-container" className="ml-2">
              <span aria-hidden="true" className="block text-xs text-gray-300"></span>
              <span aria-hidden="true" className="block font-bold text-sm">Carrinho</span>
            </div>
          </a>
        </div>
      </div>
    </div>
  );
};

export default Header;