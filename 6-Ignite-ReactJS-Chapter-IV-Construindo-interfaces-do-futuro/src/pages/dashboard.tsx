import { Header } from "../components/Header";
import { Sidebar } from "../components/SideBar";
import { Flex, SimpleGrid, Box, Text, theme } from '@chakra-ui/react';
import dynamic from "next/dynamic";

const Chart = dynamic(() => import('react-apexcharts'), {
  ssr: false,
})

const options = {
  chart: {
    toolbar: {
      show: false
    },
    zoom: {
      enabled: false,
    },
    foreColor: theme.colors.gray[500],
  },
  grid: {
    show: false
  },
  dataLabels: {
    enabled: false
  },
  stroke: {
    curve: 'smooth'
  },
  tooltip: {
    enabled: false
  },
  xaxis: {
    type: 'datetime',
    axiosBorder: {
      color: theme.colors.gray[600]
    },
    axiesTicks: {
      color: theme.colors.gray[600]
    },
    categories: [
      '2021-03-18T00:00:00.000Z',
      '2021-03-19T00:00:00.000Z',
      '2021-03-20T00:00:00.000Z',
      '2021-03-21T00:00:00.000Z',
      '2021-03-22T00:00:00.000Z',
      '2021-03-23T00:00:00.000Z',
      '2021-03-24T00:00:00.000Z',
    ]
  },
  fill: {
    opacity: 0.3,
    type: 'gradient',
    gradient: {
      shade: 'dark',
      opacityFrom: 0.7,
      opacityTo: 0.3
    }
  },
};

const series = [
  {
    name: 'series1', data: [31, 120, 10, 28, 61, 18, 109]
  }
]

export default function Dashboard() {
  return (
    <Flex direction='column' height='100vh'>

      <Header />

      <Flex width="100%" marginY="6" maxWidth={1480} marginX='auto' paddingX="6">
        <Sidebar />
        <SimpleGrid flex='1' gap='4' minChildWidth="320px" alignItems='flex-start'>
          <Box padding={['6', '8']} bgColor='gray.800' borderRadius={8} paddingBottom='4' >
            <Text fontSize='lg' marginBottom='4'>Incritos da Semana</Text>
            <Chart type="area" height={160} options={options} series={series} />
          </Box>
          <Box padding={['6', '8']} bgColor='gray.800' borderRadius={8} paddingBottom='4' >
            <Text fontSize='lg' marginBottom='4'>Taxa de Abertura</Text>
            <Chart type="area" height={160} options={options} series={series} />
          </Box>
        </SimpleGrid>
      </Flex>

    </Flex>

  )
}