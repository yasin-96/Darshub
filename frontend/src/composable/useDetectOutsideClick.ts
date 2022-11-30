import { onBeforeUnmount, onMounted } from "vue";

export default function useDetectOutsideClick(
  component: any,
  callback: Function
) {
  if (!component) return;
  if (!callback)  return;
  const listener = (event: Event) => {
    console.log(event.target, component.value);
    console.log(event.target !== component.value)
    console.log(event.composedPath().includes(component.value))

    if (
      !event.target.closest() !== component.value &&
      event.composedPath().includes(component.value)
    ) {
      return;
    }
    if (typeof callback === "function") {
      callback();
    }
  };
  onMounted(() => {
    window.addEventListener("click", listener);
  });
  onBeforeUnmount(() => {
    window.removeEventListener("click", listener);
  });

  return { listener };
}
