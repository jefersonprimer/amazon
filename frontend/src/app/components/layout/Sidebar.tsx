import React from 'react';
import SidebarSubmenu from './SidebarSubmenu';

interface SidebarProps {
  open: boolean;
  onClose: () => void;
  onOpen?: () => void;
  onOpenSubmenu?: (submenu: string) => void;
}

const Sidebar: React.FC<SidebarProps> = ({ open, onClose, onOpen, onOpenSubmenu }) => {
  // Opções de submenu para reutilização
  const submenuItems = [
    'Amazon Fire TV',
    'Amazon Music',
    'Prime Video',
    'Aplicativos Amazon',
    'Dispositivos Kindle e eBooks',
    'Echo e Alexa',
    'Audiolivros Audible',
  ];

  return (
    <div
      className={`fixed inset-0 z-40 transition-all duration-300 ${open ? 'visible' : 'invisible pointer-events-none'}`}
      aria-hidden={!open}
    >
      {/* Overlay */}
      <div
        className={`fixed inset-0  bg-opacity-40 transition-opacity duration-300 ${open ? 'opacity-100' : 'opacity-0'}`}
        onClick={onClose}
      />
      {/* Sidebar */}
      <aside
        className={`fixed top-0 left-0 h-full w-91 bg-white shadow-lg z-50 transform transition-transform duration-300 ${open ? 'translate-x-0' : '-translate-x-full'}`}
        role="dialog"
        aria-modal="true"
      >
        <button
          className="absolute top-4 -right-10 text-gray-700 hover:text-black text-2xl cursor-pointer"
          onClick={onClose}
          aria-label="Fechar menu"
        >
          <svg 
            className="text-[#F6F6F6] w-8 h-8" 
            xmlns="http://www.w3.org/2000/svg" 
            viewBox="0 0 24 24" 
            data-t="cross-svg" 
            aria-hidden="true" 
            role="img"
            fill="currentColor"
            >
                <path d="M13.414 12l5.293-5.293a.999.999 0 1 0-1.414-1.414L12 10.586 6.707 5.293a.999.999 0 1 0-1.414 1.414L10.586 12l-5.293 5.293a.999.999 0 0 0 0 1.414.993.993 0 0 0 1.414 0L12 13.414l5.293 5.293a.999.999 0 1 0 1.414-1.414L13.414 12z"></path>
          </svg>
        </button>
        <div className="p-6">
          <h2 className="text-lg font-bold mb-4 text-black">Menu</h2>
          {/* Add sidebar content here */}
          <ul className="space-y-2 border-b border-[#D5DBDB]">
            <h3>Destaques</h3>
            <li><a href="#" className="block text-gray-800 hover:text-black">Mais Vendidos</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Novidades na Amazon</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Produtos em alta</a></li>
          </ul>
          <ul className="space-y-2 border-b border-[#D5DBDB]">
            <h3>Conteúdo digital e dispositivos</h3>
            {submenuItems.map(item => (
              <li key={item} className="flex items-center justify-between cursor-pointer" onClick={() => onOpenSubmenu && onOpenSubmenu(item)}>
                <span className="block text-gray-800 hover:text-black">{item}</span>
                <span>
                  <svg 
                    className="text-[#939FA0] hover:text-black w-5 h-5" 
                    xmlns="http://www.w3.org/2000/svg" 
                    viewBox="0 0 24 24" 
                    data-t="angle-right-svg" 
                    aria-hidden="true" 
                    role="img"
                    fill="currentColor"
                  >
                    <path d="M8.6 7.4L10 6l6 6-6 6-1.4-1.4 4.6-4.6z"></path>
                  </svg>
                </span>
              </li>
            ))}
          </ul>
          <ul className="space-y-2 border-b border-[#D5DBDB]">
            <h3>Comprar por categoria</h3>
            <li><a href="#" className="block text-gray-800 hover:text-black">Amazon Fire TV</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Amazon Music</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Prime Video</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Aplicativos Amazon</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Dispositivos Kindle e eBooks</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Echo e Alexa</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Audiolivros Audible</a></li>
          </ul>
          <ul className="space-y-2 border-b border-[#D5DBDB]">
            <h3>PROGRAMAS E RECURSOS</h3>
            <li><a href="#" className="block text-gray-800 hover:text-black">Amazon Fire TV</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Amazon Music</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Prime Video</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Aplicativos Amazon</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Dispositivos Kindle e eBooks</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Echo e Alexa</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Audiolivros Audible</a></li>
          </ul>
          <ul className="space-y-2 border-b border-[#D5DBDB]">
            <h3>AJUDA E CONFIGURAÇÕES</h3>
            <li><a href="#" className="block text-gray-800 hover:text-black">Amazon Fire TV</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Amazon Music</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Prime Video</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Aplicativos Amazon</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Dispositivos Kindle e eBooks</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Echo e Alexa</a></li>
            <li><a href="#" className="block text-gray-800 hover:text-black">Audiolivros Audible</a></li>
          </ul>
        </div>
      </aside>
    </div>
  );
};

export default Sidebar; 