export default function debounce<T extends (...args: any[]) => any>(func: T, wait: number) {
    let timeout: number;
    return function(this: any, ...args: Parameters<T>) {
      const context = this;
      clearTimeout(timeout);
      timeout = setTimeout(() => {
        func.apply(context, args);
      }, wait) as any;
    } as (...args: Parameters<T>) => ReturnType<T>;
  }