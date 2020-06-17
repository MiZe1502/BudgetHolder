import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font316Black } from "../../styles/fonts";
export { FlexHorCenter, Overflowed } from "../../styles/positioning";

export const UrlWrapper = css`
    max-width: 100%;
`;

export const Url = css`
    color: ${colors.orangePrimary};
    font-weight: 500;
    
    &:hover {
        color: ${colors.orangeActive};
    }

    &:active {
        color: ${colors.orangeLight};

    }
`;