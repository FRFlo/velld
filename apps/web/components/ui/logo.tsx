import { Circle } from 'lucide-react';
import React from 'react';

const Logo: React.FC = () => {
  return (
    <div className="flex items-center gap-3 mb-12">
      <div className="h-9 w-9 rounded-xl bg-gradient-to-br from-indigo-500 to-blue-500 flex items-center justify-center shadow-sm ring-1 ring-black/[0.03]">
        <Circle className="h-4 w-4 text-white" />
      </div>
      <span className="text-base font-semibold bg-clip-text text-transparent bg-gradient-to-r from-indigo-500 to-blue-500">
        Velld
      </span>
    </div>
  );
};

export default Logo;