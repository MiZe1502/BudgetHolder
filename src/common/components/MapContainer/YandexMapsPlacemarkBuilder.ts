import { Coordinates } from "./MapsBuilder";
import { PlacemarkBuilder } from "./PlacemarkBuilder";

declare var ymaps;

export class YandexMapsPlacemarkBuilder implements PlacemarkBuilder {
    config: Record<any, any> | null = null;

    constructor(config: Record<any, any>) {
        this.config = config;
    }

    createPlacemark(coords: Coordinates) {
        return ymaps.Placemark(coords, {
            // Зададим содержимое заголовка балуна.
            balloonContentHeader: `${this.config.name}`,
            // Зададим содержимое основной части балуна.
            balloonContentBody: `${this.config.address}`,
            // Зададим содержимое нижней части балуна.
            balloonContentFooter: '<i>Footer</i>',
            // Зададим содержимое всплывающей подсказки.
            hintContent: `${this.config.name}`
        });
    }
}
