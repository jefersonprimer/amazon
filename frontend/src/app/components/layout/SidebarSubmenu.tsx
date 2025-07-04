import React from 'react';

interface SidebarSubmenuProps {
  open: boolean;
  title: string;
  options: { label: string; href?: string }[];
  onClose: () => void;
  onBack: () => void;
}

const SidebarSubmenu: React.FC<SidebarSubmenuProps> = ({ open, title, options, onClose, onBack }) => {
  return (
    <div
      className={`fixed top-0 left-0 h-full w-64 bg-white shadow-lg z-50 transform transition-transform duration-300 ${open ? 'translate-x-0' : '-translate-x-full'}`}
      role="dialog"
      aria-modal="true"
      style={{ transitionProperty: 'transform' }}
    >
      <button
        className="absolute top-4 right-4 text-gray-700 hover:text-black text-2xl"
        onClick={onClose}
        aria-label="Fechar submenu"
      >
        &times;
      </button>
      <div className="p-6">
        <button
          className="mb-4 text-blue-600 hover:underline font-medium"
          onClick={onBack}
        >
          ← MENU PRINCIPAL
        </button>
        <h2 className="text-lg font-bold mb-4 text-black">{title}</h2>
        <ul className="space-y-2">
          {options.map((option, idx) => (
            <li key={idx}>
              <a href={option.href || '#'} className="block text-gray-800 hover:text-black">
                {option.label}
              </a>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default SidebarSubmenu; 