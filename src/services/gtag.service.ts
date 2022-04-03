import { GA_MEASUREMENT_ID } from "@config/environment";

export type GTagEvent = {
  action: string;
  category?: string;
  label?: string;
  value?: number;
};

const GTagService = {
  // https://developers.google.com/analytics/devguides/collection/gtagjs/pages
  pageView(url: URL): void {
    window.gtag(`config`, GA_MEASUREMENT_ID, {
      page_path: url,
    });
  },
  // https://developers.google.com/analytics/devguides/collection/gtagjs/events
  event({ action, category, label, value }: GTagEvent): void {
    window.gtag(`event`, action, {
      event_category: category,
      event_label: label,
      value,
    });
  },
};

export default GTagService;