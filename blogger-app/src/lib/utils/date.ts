const convertToUTCAndFormat = (
	timestamp: string,
	locale: string = 'en-US',
	options: Partial<Intl.DateTimeFormatOptions> = {}
): string => {
	const date = new Date(timestamp);
	const defaultOptions: Intl.DateTimeFormatOptions = {
		year: 'numeric',
		month: 'long',
		day: 'numeric',
		hour: 'numeric',
		minute: '2-digit',
		hour12: true, // Use 12-hour clock (AM/PM)
		timeZone: 'UTC' // Convert to UTC
		// Removed timeZoneName to exclude "UTC" from output
	};
	const formatter = new Intl.DateTimeFormat(locale, {
		...defaultOptions,
		...options
	});
	return formatter.format(date);
};

export { convertToUTCAndFormat };
