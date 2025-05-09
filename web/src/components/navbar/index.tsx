"use client";

import { Button } from "@/components/ui/button";
import Logo from "./logo";
import Nav from "./nav";
import ThemeSwitcher from "./theme-switcher";
import DropDownMenu from "./drop-menu";
import clsx from "clsx";

export default function NavBar() {
  return (
    <header
      className={clsx(
        "z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 transition-transform duration-300"
      )}
    >
      <div className="mx-auto max-w-7xl px-16 bg-transparent/70 backdrop-blur-sm dash-border">
        <div className="flex h-14 items-center gap-2 md:gap-4">
          <div className="flex h-full w-full items-center justify-between px-4">
            <div className="mr-4 hidden md:flex">
              <div className="flex items-center gap-10">
                <Logo />
                <Nav />
              </div>
            </div>
            <div className="md:hidden">
              <DropDownMenu />
            </div>
            <div className="ml-auto flex items-center gap-2 md:flex-1 md:justify-end">
              <Button
                variant="secondary"
                size="icon"
                className="hover:bg-gray-100 dark:hover:bg-gray-800 cursor-pointer"
                onClick={() =>
                  window.open("https://github.com/codepzj/Stellux", "_blank")
                }
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4" />
                  <path d="M9 18c-4.51 2-5-2-7-2" />
                </svg>
              </Button>
              <ThemeSwitcher />
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}
