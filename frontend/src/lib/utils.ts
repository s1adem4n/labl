export const shouldCloseDialog = (
	e: MouseEvent & {
		currentTarget: HTMLElement;
	}
) => {
	e.stopPropagation();

	if (
		e.currentTarget.tagName !== 'DIALOG' ||
		(e.target instanceof HTMLElement && e.target.tagName !== 'DIALOG')
	)
		return false;

	const rect = e.currentTarget.getBoundingClientRect();

	const isInDialog =
		e.clientX >= rect.left &&
		e.clientX <= rect.right &&
		e.clientY >= rect.top &&
		e.clientY <= rect.bottom;

	if (!isInDialog) {
		return true;
	}
};
