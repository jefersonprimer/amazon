"use client"

import React, { useState } from 'react';
import Sidebar from './Sidebar';
import SidebarSubmenu from './SidebarSubmenu';

// Submenu options (deve ser igual ao Sidebar)
const submenuOptions: Record<string, { title: string; options: { label: string; href: string }[] }> = {
  'Amazon Fire TV': {
    title: 'Amazon Fire TV',
    options: [
      { label: 'Fire TV Stick HD', href: '/firetv/opcao1' },
      { label: 'Apps e Jogos para Fire TV Stick', href: '/firetv/opcao2' },
      { label: 'Prime Video', href: '/firetv/opcao2' }
    ],
  },
  'Amazon Music': {
    title: 'Amazon Music',
    options: [
      { label: 'Amazon Music Unlimited', href: '/music/opcao1' },
      { label: 'Amazon Music Unlimited Family', href: '/music/opcao2' },
    ],
  },
  'Prime Video': {
    title: 'Prime Video',
    options: [
      { label: 'Opção 1', href: '/primevideo/opcao1' },
      { label: 'Opção 2', href: '/primevideo/opcao2' },
    ],
  },
  'Aplicativos Amazon': {
    title: 'Aplicativos Amazon',
    options: [
      { label: 'Opção 1', href: '/aplicativos/opcao1' },
      { label: 'Opção 2', href: '/aplicativos/opcao2' },
    ],
  },
  'Dispositivos Kindle e eBooks': {
    title: 'Dispositivos Kindle e eBooks',
    options: [
      { label: 'Opção 1', href: '/kindle/opcao1' },
      { label: 'Opção 2', href: '/kindle/opcao2' },
    ],
  },
  'Echo e Alexa': {
    title: 'Echo e Alexa',
    options: [
      { label: 'Opção 1', href: '/echo/opcao1' },
      { label: 'Opção 2', href: '/echo/opcao2' },
    ],
  },
  'Audiolivros Audible': {
    title: 'Audiolivros Audible',
    options: [
      { label: 'Opção 1', href: '/audible/opcao1' },
      { label: 'Opção 2', href: '/audible/opcao2' },
    ],
  },
};

const HeaderMain = () => {
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const [activeSubmenu, setActiveSubmenu] = useState<string | null>(null);
  const [submenuOpen, setSubmenuOpen] = useState(false);

  // Quando Sidebar pedir para abrir submenu
  const handleOpenSubmenu = (submenu: string) => {
    setSidebarOpen(false);
    setActiveSubmenu(submenu);
    setSubmenuOpen(true);
  };

  // Quando clicar em voltar no submenu
  const handleBackToSidebar = () => {
    setSubmenuOpen(false);
    setActiveSubmenu(null);
    setSidebarOpen(true);
  };

  return (
    <header id="nav-main" className="relative h-12 bg-[#232F3E] text-white flex items-center px-4">
      <Sidebar
        open={sidebarOpen}
        onClose={() => setSidebarOpen(false)}
        onOpen={() => setSidebarOpen(true)}
        onOpenSubmenu={handleOpenSubmenu}
      />
      {submenuOpen && (
        <SidebarSubmenu
          open={submenuOpen}
          title={activeSubmenu ? submenuOptions[activeSubmenu]?.title || '' : ''}
          options={activeSubmenu ? submenuOptions[activeSubmenu]?.options || [] : []}
          onClose={handleBackToSidebar}
          onBack={handleBackToSidebar}
        />
      )}
      <div className="flex items-center">
        {/* Hamburger Menu */}
        <button
          id="nav-hamburger-menu"
          role="button"
          aria-label="Abrir o menu Todas as categorias"
          aria-expanded={sidebarOpen}
          data-csa-c-type="widget"
          data-csa-c-slot-id="HamburgerMenuDesktop"
          data-csa-c-interaction-events="click"
          data-csa-c-id="j1chnc-w5rhps-5ga306-phuiru"
          className="flex items-center cursor-pointer text-white no-underline hover:text-white   justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2"
          onClick={() => setSidebarOpen(true)}
        >
          <span>
            <svg 
                className="text-white w-5 h-5" 
                xmlns="http://www.w3.org/2000/svg" 
                viewBox="0 0 24 24" 
                data-t="menu-svg" 
                aria-hidden="true" 
                role="img"
                fill="currentColor"
            >
                <path className="menu-ab-svg__top--Xgjtj" d="M20 5H4V7H20V5Z"></path>
                <path className="menu-ab-svg__middle--vsB5D" d="M4 11H20V13H4V11Z"></path>
                <path className="menu-ab-svg__bottom--6yklx" d="M20 17H4V19H20V17Z"></path>
            </svg>
          </span>
          <span className="hm-icon-label">Todos</span>
        </button>
        <div id="nav-shop"></div>
      </div>
      
      <div className="flex-grow">
        <div id="nav-xshop-container">
          <div id="nav-xshop" className="nav-progressive-content">
            <ul className="flex space-x-2 px-4">
              <li className="nav-li  flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/gp/browse.html?node=17877554011&amp;ld=ASBRSOA_retail_sell_header_t1&amp;ref_=nav_cs_sell" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_0" data-csa-c-content-id="nav_cs_sell" data-csa-c-id="l2h69k-w0ixl-h0c5ii-h5shs7">
                    Venda na Amazon
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/gp/bestsellers/?ref_=nav_cs_bestsellers" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_1" data-csa-c-content-id="nav_cs_bestsellers" data-csa-c-id="2st444-y0cdr0-4plw72-gs2teb">
                    Mais Vendidos
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/deals?ref_=nav_cs_gb" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_2" data-csa-c-content-id="nav_cs_gb" data-csa-c-id="nko9o2-pewuc9-1rzv8t-svmoh7">
                    Ofertas do Dia
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div" id="nav-link-amazonprime">
                  <a href="/prime?ref_=nav_cs_primelink_nonmember" className="nav-a flex items-center justify-center text-white text-sm no-underline hover:text-white" data-ux-jq-mouseenter="true" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav-link-amazonprime" data-csa-c-content-id="nav_cs_primelink_nonmember" data-csa-c-id="b4b15a-grgx6u-llorbg-g86bo7">
                    <span>Prime</span>
                    <span>
                    <svg
                        className="text-[#666666] w-5 h-5"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        data-t="dropdown-svg"
                        aria-hidden="true"
                        role="img"
                        fill='currentColor'
                    >
                        <path d="M7 10h10l-5 5z"></path>
                    </svg>
                    </span>
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/Livros/b/?ie=UTF8&amp;node=6740748011&amp;ref_=nav_cs_books" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_4" data-csa-c-content-id="nav_cs_books" data-csa-c-id="dcm1sf-1z7sbe-17ua92-5gps0z">
                    Livros
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/gp/new-releases/?ref_=nav_cs_newreleases" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_5" data-csa-c-content-id="nav_cs_newreleases" data-csa-c-id="xm416h-d4tca7-epx7eu-fkyl6k">
                    Novidades na Amazon
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/music/unlimited?ref_=nav_cs_music" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_6" data-csa-c-content-id="nav_cs_music" data-csa-c-id="vw2tc1-6ripid-7sg6j6-zdkgdz">
                    Música
                  </a>
                </div>
              </li>
              <li className="nav-li flex items-center justify-center hover:outline-2 hover:outline-white transition-all h-10 px-2">
                <div className="nav-div">
                  <a href="/Games-Consoles/b/?ie=UTF8&amp;node=7791985011&amp;ref_=nav_cs_video_games" className="nav-a text-white text-sm no-underline hover:text-white" tabIndex={0} data-csa-c-type="link" data-csa-c-slot-id="nav_cs_7" data-csa-c-content-id="nav_cs_video_games" data-csa-c-id="cwvyjg-iki6l5-79pwby-numcd6">
                    Games
                  </a>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <div className="flex-shrink-0 ml-auto">
        {/* Navyaan SWM */}
        <div id="nav-swmslot">
          <div id="navSwmHoliday" className="h-10 w-auto overflow-hidden relative" data-cel-widget="nav_sitewide_msg">
            <a aria-label="X-Site" href="/gp/kindle/ku/sign-up/?ie=UTF8&amp;ie=UTF8&amp;ref_=nav_swm_ku_swm_081170&amp;pf_rd_p=b3607782-b696-4dd4-ad6b-afa62e1cff27&amp;pf_rd_s=nav-sitewide-msg&amp;pf_rd_t=4201&amp;pf_rd_i=navbar-4201&amp;pf_rd_m=A2Q3Y263D00KWC&amp;pf_rd_r=2JXJAEWYPJHGMRE06RV2" className="nav-imageHref" tabIndex={0}>
              <img alt="X-Site" src="https://m.media-amazon.com/images/G/32/kindle/ku/2017/SWM/SWM400x39v1._CB485941404_.png" className="h-full w-full object-cover"/>
            </a>
          </div>
        </div>
      </div>
    </header>
  );
};

export default HeaderMain;