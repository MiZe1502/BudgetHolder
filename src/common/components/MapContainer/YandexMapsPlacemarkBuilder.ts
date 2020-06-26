import { Coordinates } from "./MapsBuilder";
import { PlacemarkBuilder } from "./PlacemarkBuilder";
import { MapItemData } from "../MapActionElement/utils";

declare var ymaps;

export class YandexMapsPlacemarkBuilder implements PlacemarkBuilder {

    createPlacemark(coords: Coordinates, data: MapItemData) {
        //return new ymaps.Placemark(coords);
        return new ymaps.Placemark(coords, {
            // Зададим содержимое заголовка балуна.
            balloonContentHeader: `${data.name}`,
            // Зададим содержимое основной части балуна.
            balloonContentBody: `${data.address}`,
            // Зададим содержимое нижней части балуна.
            balloonContentFooter: '<i>Footer</i>',
            // Зададим содержимое всплывающей подсказки.
            hintContent: `${data.name}`
        });
    }
}
