import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import rehypeRaw from "rehype-raw";
import rehypeHighlight from "rehype-highlight";
import "@/styles/md.scss";
import { Terminal } from "lucide-react";
import CopyButton from "@/components/markdown/components/copy";
import React from "react";

export default function MdSimple({ content }: { content: string }) {
  let index = 1;
  return (
    <div className="markdown-body">
      <Markdown
        rehypePlugins={[rehypeRaw, rehypeHighlight]}
        remarkPlugins={[remarkGfm]}
        components={{
          h2: ({ children }) => {
            return <h2 id={`header-${index++}`}>{children}</h2>;
          },
          h3: ({ children }) => {
            return <h3 id={`header-${index++}`}>{children}</h3>;
          },
          h4: ({ children }) => {
            return <h4 id={`header-${index++}`}>{children}</h4>;
          },
          h5: ({ children }) => {
            return <h5 id={`header-${index++}`}>{children}</h5>;
          },
          h6: ({ children }) => {
            return <h6 id={`header-${index++}`}>{children}</h6>;
          },
          pre: ({ children }) => <pre style={{ padding: "0" }}>{children}</pre>,
          code: ({ node, className, children, ...props }) => {
            // 判断是否为代码块
            const match = /language-(\w+)/.exec(className || "");
            const count =
              React.Children.toArray(children).length === 1
                ? (
                    React.Children.toArray(children)[0]
                      .toString()
                      .match(/\n/g) || []
                  ).length
                : 0;

            if (match?.length || count > 0) {
              // 代码块生成唯一id
              const id = Math.random().toString(36).slice(2, 11);
              return (
                <div className="not-prose rounded-md">
                  <div className="flex h-10 items-center justify-between px-4 bg-zinc-100 dark:bg-zinc-800">
                    <div className="flex items-center gap-2">
                      <Terminal size={16} />
                      <span className="h-full flex items-center">
                        {node?.data?.meta || match?.[1] || "plaintext"}
                      </span>
                    </div>
                    <CopyButton id={id} />
                  </div>
                  <div className="overflow-x-auto">
                    <div id={id} className="p-4">
                      {children}
                    </div>
                  </div>
                </div>
              );
            }
            return (
              <code {...props} className="not-prose rounded-md">
                {children}
              </code>
            );
          },
        }}
      >
        {content}
      </Markdown>
    </div>
  );
}
